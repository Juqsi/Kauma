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
	return map[string]interface{}{"S": args.Z}, nil
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
func handleGfpolyDiv(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyDiv)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"Q": args.Q, "R": args.R}, nil
}
func handleGfpolyPowmod(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyPowmod)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"Z": args.Z}, nil
}
func handleGfpolySort(argsData []byte) (map[string]interface{}, error) {
	var args gfpoly.GfpolySort
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"sorted_polys": args.SortedPolys}, nil
}
func handleGfpolyMakeMonic(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyMakeMonic)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"A*": args.ASternchen}, nil
}

func handleGfpolySqrt(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolySqrt)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"S": args.S}, nil
}

func handleGfpolyDiff(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyDiff)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"F'": args.FStrich}, nil
}

func handleGfpolyGgt(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolyGgt)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"G": args.G}, nil
}

func handleGfpolySff(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolySff)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"factors": args.Factors}, nil
}

func handleGfpolyDdf(argsData []byte) (map[string]interface{}, error) {
	args := new(gfpoly.GfpolySff)
	if err := json.Unmarshal(argsData, &args); err != nil {
		return nil, err
	}
	args.Execute()
	return map[string]interface{}{"factors": args.Factors}, nil
}
