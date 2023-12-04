package main

import "fmt"

func main() {
	fmt.Println(areOccurrencesEqual("abacbc")) // true
	fmt.Println(areOccurrencesEqual("aaabb"))  // false
}

func areOccurrencesEqual(s string) bool {
	sM := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sM[s[i]]++
	}

	firstElement := sM[s[0]]

	for _, v := range sM {
		if v == firstElement {
			continue
		} else {
			return false
		}
	}
	return true
}
