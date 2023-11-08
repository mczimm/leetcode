package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))         //true
	fmt.Println(repeatedSubstringPattern("aba"))          //false
	fmt.Println(repeatedSubstringPattern("abcabcabcabc")) //true
	fmt.Println(repeatedSubstringPattern("abac"))         //false
	fmt.Println(repeatedSubstringPattern("ababab"))       //true
}

// From the solution: https://leetcode.com/problems/repeated-substring-pattern/solutions/3938580/99-42-2-approaches-o-n/
func repeatedSubstringPattern(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}
