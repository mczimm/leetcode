package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))                //"dc-ba"
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))        //"j-Ih-gfE-dCba"
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) //"Qedo1ct-eeLg=ntse-T!"
}

func reverseOnlyLetters(s string) string {
	var res = make([]byte, len(s))
	var left, right = 0, len(s) - 1

	for left <= right {
		switch {
		case unicode.IsLetter(rune(s[left])) && unicode.IsLetter(rune(s[right])):
			res[left] = s[right]
			res[right] = s[left]
			left++
			right--
		case !unicode.IsLetter(rune(s[left])):
			res[left] = s[left]
			left++
		case !unicode.IsLetter(rune(s[right])):
			res[right] = s[right]
			right--
		}
	}

	return string(res)
}
