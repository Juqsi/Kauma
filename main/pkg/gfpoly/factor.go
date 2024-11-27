package gfpoly

import (
	"sort"
)

func (list *Factors) Sort() *Factors {
	sort.SliceStable(*list, func(i, j int) bool {
		return (*list)[i].Factor.Cmp(&(*list)[i].Factor, &(*list)[j].Factor) < 0
	})
	return list
}

type Factor struct {
	Factor   Poly
	Exponent int
}

type Factors []Factor

type FactorModel struct {
	Factor   []string `json:"factor"`
	Exponent int      `json:"exponent"`
}
type FactorsModel []FactorModel

func (list Factors) Base64() FactorsModel {
	factorsModel := make(FactorsModel, len(list))
	for i, f := range list {
		factorsModel[i].Factor = f.Factor.Base64()
		factorsModel[i].Exponent = f.Exponent
	}
	return factorsModel
}

type FactorsModelWithDegree []FactorModelWithDegree

type FactorModelWithDegree struct {
	Factor []string `json:"factor"`
	Degree int      `json:"degree"`
}

func (list Factors) Base64Degree() FactorsModelWithDegree {
	factorsModel := make(FactorsModelWithDegree, len(list))
	for i, f := range list {
		factorsModel[i].Factor = f.Factor.Base64()
		factorsModel[i].Degree = f.Exponent
	}
	return factorsModel
}
