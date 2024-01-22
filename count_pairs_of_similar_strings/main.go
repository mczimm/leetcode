package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
}

func similarPairs(words []string) int {
	var res int
	//wordsHM := make([]map[byte]struct{}, len(words))
	wordsSl := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		tm := make(map[byte]struct{})
		for j := 0; j < len(words[i]); j++ {
			tm[words[i][j]] = struct{}{}
		}

		var r string
		for ch := range tm {
			r += string(ch)
		}
		wordsSl[i] = sortString(r)
	}

	for i := 0; i < len(wordsSl); i++ {
		for j := i + 1; j < len(wordsSl); j++ {
			if wordsSl[i] == wordsSl[j] {
				res++
			}
		}
	}

	return res
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
