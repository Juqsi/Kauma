package testcases

import (
	"Abgabe/main/pkg/gfpoly"
	"encoding/json"
)

func handleGfpolyAdd(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyAdd)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"Z": args.Z}, nil
}

func handleGfpolyMul(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyMul)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"P": args.P}, nil
}

func handleGfpolyPow(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyPow)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"Z": args.Z}, nil
}
