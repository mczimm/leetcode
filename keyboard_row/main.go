package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"})) // ["Alaska","Dad"]
	fmt.Println(findWords([]string{"omk"}))                             // []
	fmt.Println(findWords([]string{"adsdf", "sfd"}))                    // ["adsdf","sfd"]
}

func findWords(words []string) []string {
	first := make(map[rune]struct{})
	second := make(map[rune]struct{})
	third := make(map[rune]struct{})
	var res []string

	for _, v := range "qwertyuiop" {
		first[v] = struct{}{}
	}
	for _, v := range "asdfghjkl" {
		second[v] = struct{}{}
	}
	for _, v := range "zxcvbnm" {
		third[v] = struct{}{}
	}

	for i := 0; i < len(words); i++ {
		var tmpFirst, tmpSecond, tmpThird []rune
		for j := 0; j < len(words[i]); j++ {
			if _, ok := first[unicode.ToLower(rune(words[i][j]))]; ok {
				tmpFirst = append(tmpFirst, rune(words[i][j]))
			} else {
				tmpFirst = []rune{}
				break
			}
		}
		for j := 0; j < len(words[i]); j++ {
			if _, ok := second[unicode.ToLower(rune(words[i][j]))]; ok {
				tmpSecond = append(tmpSecond, rune(words[i][j]))
			} else {
				tmpSecond = []rune{}
				break
			}
		}
		for j := 0; j < len(words[i]); j++ {
			if _, ok := third[unicode.ToLower(rune(words[i][j]))]; ok {
				tmpThird = append(tmpThird, rune(words[i][j]))
			} else {
				tmpThird = []rune{}
				break
			}
		}
		if len(tmpFirst) > 0 {
			res = append(res, string(tmpFirst))
		}
		if len(tmpSecond) > 0 {
			res = append(res, string(tmpSecond))
		}
		if len(tmpThird) > 0 {
			res = append(res, string(tmpThird))
		}
	}

	return res
}
