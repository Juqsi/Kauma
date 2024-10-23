package utils

import "os"

func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
