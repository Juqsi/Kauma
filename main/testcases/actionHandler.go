package testcases

import (
	"Abgabe/main/pkg/actions"
	"encoding/json"
)

func handlePoly2Block(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Poly2Block)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"block": args.Result}, nil
}

func handleBlock2Poly(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Block2Poly)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"coefficients": args.Result}, nil
}

func handleGfmul(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Gfmul)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"product": args.Result}, nil
}

func handleSea128(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Sea128)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"output": args.Result}, nil
}

func handleXex(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Xex)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"output": args.Result}, nil
}

func handlePaddingOracle(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.PaddingOracle)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"plaintext": args.Result}, nil
}

func handleGcmEncrypt(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Gcm_Encrypt)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{
		"ciphertext": args.Ciphertext,
		"tag":        args.Tag,
		"L":          args.L,
		"H":          args.H,
	}, nil
}

func handleGcmDecrypt(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Gcm_Decrypt)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{
		"authentic": args.Authentic,
		"plaintext": args.Plaintext,
	}, nil
}

func handleGfdiv(argsData []byte) (map[string]interface{}, error) {
	args := new(actions.Gfdiv)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"q": args.Result}, nil
}
