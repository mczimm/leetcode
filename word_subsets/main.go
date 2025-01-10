package main

import "fmt"

func main() {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}
	fmt.Println(wordSubsets(words1, words2)) //["facebook", "google", "leetcode"]

	words1 = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 = []string{"l", "e"}
	fmt.Println(wordSubsets(words1, words2)) //["apple", "google", "leetcode"]
}

func wordSubsets(words1 []string, words2 []string) []string {
	var result []string

	// Helper function to compute frequency of characters in a word
	getFreq := func(word string) map[rune]int {
		freq := make(map[rune]int)
		for _, v := range word {
			freq[v]++
		}
		return freq
	}

	// Compute max frequency for each character in words2
	maxFreq := make(map[rune]int)
	for _, v := range words2 {
		vFreq := getFreq(v)
		for char, cnt := range vFreq {
			if cnt > maxFreq[char] {
				maxFreq[char] = cnt
			}
		}
	}

	// Check each word in words1 against maxFrequency
	for _, v := range words1 {
		vFreq := getFreq(v)
		isUniversal := true
		for char, cnt := range maxFreq {
			if vFreq[char] < cnt {
				isUniversal = false
				break
			}
		}
		if isUniversal {
			result = append(result, v)
		}
	}
	return result
}
