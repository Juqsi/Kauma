package gfpoly

import (
	"sort"
)

type GfpolySort struct {
	Polys       [][]string `json:"polys"`
	SortedPolys [][]string `json:"sorted_polys"`
}

func (g *GfpolySort) Execute() {
	g.SortedPolys = NewPolyListFromBase64(g.Polys).Sort().Base64()
}

func (list *Polys) Sort() *Polys {
	sort.SliceStable(*list, func(i, j int) bool {
		return (*list)[i].Cmp((*list)[i], (*list)[j]) < 0
	})
	return list
}
