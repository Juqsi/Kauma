package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/utils"
	"math/big"
	"sync"
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
	Ciphertext     utils.Text
	AssociatedData utils.Text
	Tag            big.Int
	L              big.Int
	Poly           Poly
}

func (args *GcmCrack) Execute() {
	messages := make([]mes, 0, 3)
	for _, m := range []message{args.M1, args.M2, args.M3} {
		me := new(mes)
		me.Ciphertext = utils.GetContent(m.Ciphertext)
		me.AssociatedData = utils.GetContent(m.AssociatedData)
		me.Tag = utils.NewBigEndianLongFromGcmInBase64(m.Tag).Int
		_, me.L = actions.CalculateL(m.Ciphertext, m.AssociatedData)
		me.Poly = *New128PolyFromFactors([]utils.Text{me.AssociatedData, me.Ciphertext, {Content: me.L, Len: 16}, {Content: me.Tag, Len: 16}})
		messages = append(messages, *me)
	}
	poly := new(Poly).Add(&messages[0].Poly, &messages[1].Poly)
	candidates := poly.FindRoots()

	for _, candidate := range candidates {
		ghash := actions.GHASHBigEndian(candidate, messages[0].Ciphertext, messages[0].L, messages[0].AssociatedData)
		mask := ghash.Xor(&ghash, &messages[0].Tag)
		ghashM3 := actions.GHASHBigEndian(candidate, messages[2].Ciphertext, messages[2].L, messages[2].AssociatedData)
		tagM3 := new(big.Int).Xor(mask, &ghashM3)
		if tagM3.Cmp(&messages[2].Tag) == 0 {
			forgeryCipherText := utils.GetContent(args.Forgery.Ciphertext)
			forgeryAd := utils.GetContent(args.Forgery.AssociatedData)
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
	candidates := make(chan big.Int, len(factors))

	var wg sync.WaitGroup
	for _, factor := range factors {
		wg.Add(1)
		go func(factor Factor) {
			defer wg.Done()
			if factor.Factor.Degree() > 1 {
				ddfFactors := factor.Factor.Ddf()
				for _, dFactor := range ddfFactors {
					if dFactor.Factor.Degree() == 1 {
						candidates <- dFactor.Factor[0]
					} else if dFactor.Exponent == 1 {
						edfFactors := new(Polys).Edf(&dFactor.Factor, dFactor.Exponent)
						for _, eFactor := range edfFactors {
							if eFactor.Degree() == 1 {
								candidates <- eFactor[0]
							}
						}
					}
				}
			}
		}(factor)
	}

	go func() {
		wg.Wait()
		close(candidates)
	}()

	var result []big.Int
	for candidate := range candidates {
		result = append(result, candidate)
	}

	return result
}
