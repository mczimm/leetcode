package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(detectCapitalUse("USA"))    //true
	fmt.Println(detectCapitalUse("flaG"))   //false
	fmt.Println(detectCapitalUse("Nebius")) //true
	fmt.Println(detectCapitalUse("FlaG"))   //false
}

// From the solutions
func detectCapitalUse(word string) bool {
	allCap := strings.ToUpper(word)
	allLower := strings.ToLower(word)
	firstCap := strings.ToUpper(string(word[0])) + strings.ToLower(word)[1:]

	return word == allCap || word == allLower || word == firstCap
}
