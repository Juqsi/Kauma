package testcases

import (
	"Abgabe/main/pkg/glasskey"
	"encoding/json"
)

func handleGlasskeyPrng(argsData []byte) (map[string]interface{}, error) {
	args := new(glasskey.Prng)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{
		"blocks": args.Blocks,
	}, nil
}
