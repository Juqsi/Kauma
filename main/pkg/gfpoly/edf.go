package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"math/big"
	"sync"
)

type GfpolyEdf struct {
	F       []string `json:"F"`
	D       int      `json:"d"`
	Factors [][]string
}

func (g *GfpolyEdf) Execute() {
	polyA := NewPolyFromBase64(g.F)
	factors := new(Polys).Edf(polyA, g.D)
	g.Factors = factors.Sort().Base64()
}

func (p *Polys) Edf(f *Poly, d int) Polys {
	q := big.NewInt(1)
	q.Lsh(q, 128)
	n := f.Degree() / d

	z := Polys{*f}

	exp := new(big.Int).Exp(q, big.NewInt(int64(d)), nil)
	exp.Sub(exp, big.NewInt(1))
	exp.Div(exp, big.NewInt(3))
	g := new(Poly)

	for len(z) < n {
		h := RandomPolynomial(f.Degree())

		g.PowMod(h, exp, f)
		g.Add(g, &Poly{actions.OneBlock})

		const maxWorkers = 4
		workerPool := make(chan struct{}, maxWorkers)

		type Result struct {
			Index int
			NewZ  []Poly
		}

		results := make(chan Result, len(z))
		var wg sync.WaitGroup

		for i := len(z) - 1; i >= 0; i-- {
			u := z[i]
			if u.Degree() > d {
				workerPool <- struct{}{} // Blockiert, wenn maxWorkers erreicht
				wg.Add(1)
				go func(i int, u Poly, g *Poly) {
					defer wg.Done()
					defer func() { <-workerPool }() // Gibt einen Worker frei

					j := new(Poly).Gcd(&u, g)
					if !j.IsOne() && j.Cmp(&u) != 0 {
						tmp, _ := new(Poly).Div(&u, j)
						results <- Result{
							Index: i,
							NewZ:  []Poly{*j.makeMonic(j), *tmp.makeMonic(tmp)},
						}
					} else {
						results <- Result{
							Index: i,
							NewZ:  []Poly{u},
						}
					}
				}(i, u, g)
			}
		}

		wg.Wait()
		close(results)
		close(workerPool)

		newZMap := make(map[int][]Poly)
		for res := range results {
			newZMap[res.Index] = res.NewZ
		}

		newZ := make([]Poly, 0, len(z)*2)
		for i := 0; i < len(z); i++ {
			if updated, ok := newZMap[i]; ok {
				newZ = append(newZ, updated...)
			} else {
				newZ = append(newZ, z[i])
			}
		}

		z = newZ
	}

	*p = z
	return *p
}
