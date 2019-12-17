package go_advent_of_code_2019

import (
	"io/ioutil"
	"strings"
)

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bytes))
}
