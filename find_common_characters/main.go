package main

import "fmt"

func main() {
	//["e","l","l"]
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	//[]
	fmt.Println(commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}))
	//["c","o"]
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func commonChars(words []string) []string {
	// Create a map to store the count of each character in the first string.
	chCNT := make(map[rune]int)
	for _, v := range words[0] {
		chCNT[v]++
	}

	// Iterate through the remaining strings and update the character count in the map.
	for _, word := range words[1:] {
		tmpCNT := make(map[rune]int)
		for _, char := range word {
			tmpCNT[char]++
		}
		for ch, count := range chCNT {
			if tmpCNT[ch] < count {
				chCNT[ch] = tmpCNT[ch]
			}
		}
	}
	var res []string
	for ch, cnt := range chCNT {
		for i := 0; i < cnt; i++ {
			res = append(res, string(ch))
		}
	}
	return res
}
