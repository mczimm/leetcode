package main

import "fmt"

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"})) //2
}

func maximumNumberOfStringPairs(words []string) int {
	var res int
	wordsH := make(map[string]struct{})

	for _, word := range words {
		tmp := string(word[1]) + string(word[0])
		if _, ok := wordsH[tmp]; ok {
			res++
		} else {
			wordsH[word] = struct{}{}
		}
	}

	return res
}
