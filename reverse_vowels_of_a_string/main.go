package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("hello"))    //holle
	fmt.Println(reverseVowels("leetcode")) //leotcede
}

func reverseVowels(s string) string {
	arr := strings.Split(s, "")
	left := 0
	right := len(s) - 1

	for i := 0; i < len(s); i++ {
		if isVowel(string(s[left])) {
			tmp := string(s[left])
			for {
				if isVowel(string(s[right])) {
					arr[left] = string(s[right])
					arr[right] = tmp
					right--
					break
				}
				right--
			}
		}
		left++
	}

	return strings.Join(arr, "")
}

func isVowel(character string) bool {
	if character == "a" || character == "e" || character == "i" || character == "o" || character == "u" {
		return true
	} else if character == "A" || character == "E" || character == "I" || character == "O" || character == "U" {
		return true
	} else {
		return false
	}
}
