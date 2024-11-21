package gfpoly

import "sync"

type GfpolySort struct {
	Polys       [][]string `json:"polys"`
	SortedPolys [][]string `json:"sorted_polys"`
}

func (g *GfpolySort) Execute() {
	g.SortedPolys = NewPolyListFromBase64(g.Polys).Sort().Base64()
}

func (polys *Polys) Sort() *Polys {
	if len(*polys) < 2 {
		return polys
	}
	left, right := 0, len(*polys)-1
	pivotIndex := len(*polys) / 2

	(*polys)[pivotIndex], (*polys)[right] = (*polys)[right], (*polys)[pivotIndex]

	for i := range *polys {
		if (*polys)[i].Degree() < (*polys)[right].Degree() {
			(*polys)[i], (*polys)[left] = (*polys)[left], (*polys)[i]
			left++
		} else if (*polys)[i].Degree() == (*polys)[right].Degree() && i != right {
			index := (*polys)[i].Degree() - 1
			for index >= 0 {
				a := (*polys)[i][index].Cmp(&(*polys)[right][index])
				if a < 0 {
					(*polys)[i], (*polys)[left] = (*polys)[left], (*polys)[i]
					left++
					break
				} else if a > 0 {
					break
				}
				index--
			}
		}
	}

	(*polys)[left], (*polys)[right] = (*polys)[right], (*polys)[left]

	leftPart := (*polys)[:left]
	rightPart := (*polys)[left+1:]

	const threshold = 1000

	if len(leftPart) > threshold && len(rightPart) > threshold {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			leftPart.Sort()
		}()

		go func() {
			defer wg.Done()
			rightPart.Sort()
		}()

		wg.Wait()
	} else {
		leftPart.Sort()
		rightPart.Sort()
	}

	return polys
}
