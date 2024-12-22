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

	const maxWorkers = 4
	workerPool := make(chan struct{}, maxWorkers)
	jobQueue := make(chan int, len(z))
	resultQueue := make(chan map[int][]Poly, len(z))

	for i := 0; i < maxWorkers; i++ {
		go func() {
			for i := range jobQueue {
				u := z[i]
				newZMap := make(map[int][]Poly)
				if u.Degree() > d {
					j := new(Poly).Gcd(&u, g)
					if !j.IsOne() && j.Cmp(&u) != 0 {
						tmp, _ := new(Poly).Div(&u, j)
						newZMap[i] = []Poly{*j.makeMonic(j), *tmp.makeMonic(tmp)}
					} else {
						newZMap[i] = []Poly{u}
					}
				}
				resultQueue <- newZMap
			}
		}()
	}

	for len(z) < n {
		h := RandomPolynomial(f.Degree())

		g.PowMod(h, exp, f)
		g.Add(g, &Poly{actions.OneBlock})

		newZMap := make(map[int][]Poly)
		var wg sync.WaitGroup

		for i := len(z) - 1; i >= 0; i-- {
			u := z[i]
			if u.Degree() > d {
				workerPool <- struct{}{} // Blockiert, wenn maxWorkers erreicht
				wg.Add(1)
				go p.processPoly(i, u, g, workerPool, newZMap, &wg)
			}
		}

		wg.Wait()

		z = p.rebuildZ(z, newZMap)
	}

	*p = z
	return *p
}

func (p *Polys) processPoly(i int, u Poly, g *Poly, workerPool chan struct{}, newZMap map[int][]Poly, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { <-workerPool }() // Worker freigeben

	j := new(Poly).Gcd(&u, g)
	if !j.IsOne() && j.Cmp(&u) != 0 {
		tmp, _ := new(Poly).Div(&u, j)
		newZMap[i] = []Poly{*j.makeMonic(j), *tmp.makeMonic(tmp)}
	} else {
		newZMap[i] = []Poly{u}
	}
}

func (p *Polys) rebuildZ(z []Poly, newZMap map[int][]Poly) []Poly {
	newZ := make([]Poly, 0, len(z)*2)
	for i := 0; i < len(z); i++ {
		if updated, ok := newZMap[i]; ok {
			newZ = append(newZ, updated...)
		} else {
			newZ = append(newZ, z[i])
		}
	}
	return newZ
}
