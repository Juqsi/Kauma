package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/utils"
	"math/big"
)

type message struct {
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Tag            string `json:"tag"`
}
type Forgery struct {
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
}
type GcmCrack struct {
	Nonce   string  `json:"nonce"`
	M1      message `json:"m1"`
	M2      message `json:"m2"`
	M3      message `json:"m3"`
	Forgery Forgery `json:"forgery"`
	Tag     string  `json:"tag"`
	H       string  `json:"H"`
	Mask    string  `json:"mask"`
}

type mes struct {
	Ciphertext     big.Int
	AssociatedData big.Int
	Tag            big.Int
	L              big.Int
	Poly           Poly
}

func (args *GcmCrack) Execute() {
	//nonce := utils.NewLongFromBase64(args.Nonce).Int

	messages := make([]mes, 0, 3)
	for _, m := range []message{args.M1, args.M2, args.M3} {
		me := new(mes)
		me.Ciphertext = utils.NewBigEndianLongFromGcmInBase64(m.Ciphertext).Int
		me.AssociatedData = utils.NewBigEndianLongFromGcmInBase64(m.AssociatedData).Int
		me.Tag = utils.NewBigEndianLongFromGcmInBase64(m.Tag).Int
		_, me.L = actions.CalculateL(m.Ciphertext, m.AssociatedData)
		me.Poly = *New128PolyFromFactors([]big.Int{me.AssociatedData, me.Ciphertext, me.L, me.Tag})
		messages = append(messages, *me)
	}
	poly := new(Poly).Add(&messages[0].Poly, &messages[1].Poly)
	//H candiandes calculation
	candidates := poly.FindRoots()
	//calculation of GHASH
	for _, candidate := range candidates {
		ghash := actions.GHASHBigEndian(candidate, messages[0].Ciphertext, messages[0].L, messages[0].AssociatedData)
		mask := ghash.Xor(&ghash, &messages[0].Tag)
		ghashM3 := actions.GHASHBigEndian(candidate, messages[2].Ciphertext, messages[2].L, messages[2].AssociatedData)
		tagM3 := new(big.Int).Xor(mask, &ghashM3)
		if tagM3.Cmp(&messages[2].Tag) == 0 {
			forgeryCipherText := utils.NewBigEndianLongFromGcmInBase64(args.Forgery.Ciphertext).Int
			forgeryAd := utils.NewBigEndianLongFromGcmInBase64(args.Forgery.AssociatedData).Int

			_, forgeryL := actions.CalculateL(args.Forgery.Ciphertext, args.Forgery.AssociatedData)
			resultGhash := actions.GHASHBigEndian(candidate, forgeryCipherText, forgeryL, forgeryAd)

			forgeryTag := *new(big.Int).Xor(&resultGhash, mask)

			args.Tag = utils.NewLongFromBigInt(forgeryTag).GcmToggle().GetBase64(16)
			args.H = utils.NewLongFromBigInt(candidate).GcmToggle().GetBase64(16)
			args.Mask = utils.NewLongFromBigInt(*mask).GcmToggle().GetBase64(16)
			return
		}
	}

	return

}

func (p *Poly) FindRoots() []big.Int {
	factors := p.sff()
	candidates := []big.Int{}

	for _, factor := range factors {
		if factor.Factor.Degree() > 1 {
			ddfFactors := factor.Factor.Ddf()
			for _, dFactor := range ddfFactors {
				if dFactor.Factor.Degree() == 1 {
					candidates = append(candidates, dFactor.Factor[0])
				} else if dFactor.Exponent == 1 {
					edfFactors := new(Polys).Edf(&dFactor.Factor, dFactor.Exponent)
					for _, eFactor := range edfFactors {
						if eFactor.Degree() == 1 {
							candidates = append(candidates, eFactor[0])
						}
					}
				}
			}
		}
	}
	return candidates
}
