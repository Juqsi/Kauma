package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"math/big"
)

type GfpolyPowMod struct {
	A []string `json:"A"`
	M []string `json:"M"`
	K big.Int  `json:"K"`
	Z []string `json:"p"`
}

func (g *GfpolyPowMod) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyM := NewPolyFromBase64(g.M)
	polyA.PowMod(polyA, g.K, polyM)
	g.Z = polyA.Base64()
}

func (p *Poly) PowMod(base *Poly, k big.Int, m *Poly) *Poly {
	result := &Poly{actions.OneBlock}

	workingBase := base.DeepCopy()

	for k.Sign() != 0 {
		if k.Bit(0) == 1 {
			result.Mul(result, workingBase)
			result.Mod(result, m)
		}

		workingBase.Mul(workingBase, workingBase)
		workingBase.Mod(workingBase, m)

		k.Rsh(&k, 1)
	}
	*p = *result.CutLeadingZeroFaktors()
	return p
}
