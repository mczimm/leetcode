package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

var result []string

func letterCasePermutation(s string) []string {
	res := traverse("", s, 0)
	result = []string{}
	return res
}

func traverse(current string, s string, i int) []string {

	if len(current) == len(s) {
		result = append(result, current)
		return result
	}

	traverse(current+string(s[i]), s, i+1)
	if unicode.IsLetter(rune(s[i])) {
		swap := swapCase(string(s[i]))
		traverse(current+swap, s, i+1)
	}

	return result
}

func swapCase(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case unicode.IsLower(r):
			return unicode.ToUpper(r)
		case unicode.IsUpper(r):
			return unicode.ToLower(r)
		}
		return r
	}, s)
}
