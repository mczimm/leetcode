package main

import "strings"

func main() {
	println(stringMatching([]string{"mass", "as", "hero", "superhero"}))             // ["as","hero"]
	println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"})) // ["leetcode","od","am"]
}

func stringMatching(words []string) []string {
	var ans []string
	tmp := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if words[i] != words[j] {
				if strings.Contains(words[i], words[j]) {
					//ans = append(ans, words[j])
					tmp[words[j]] = true
				}
			}
		}
	}
	for i := range tmp {
		ans = append(ans, i)
	}
	return ans
}
