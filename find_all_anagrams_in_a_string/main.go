package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("baa", "aa"))         //[1]
}

// My attempt: Time Limit Exceeded
func findAnagrams(s string, p string) []int {
	var res []int
	mapP := make(map[byte]int)
	mapTMP := make(map[byte]int)

	if len(s) < len(p) {
		return res
	}

	for i := 0; i < len(p); i++ {
		mapP[p[i]]++
	}

	for i := 0; i < len(s); i++ {
		for k, v := range mapP {
			mapTMP[k] = v
		}
		pointer := i
		exists := true
		if len(s)-pointer >= len(p) {
			for j := 0; j < len(p); j++ {
				if mapTMP[s[pointer]] == 0 {
					delete(mapTMP, s[pointer])
				}
				if _, ok := mapTMP[s[pointer]]; !ok {
					exists = false
					break
				} else {
					mapTMP[s[pointer]]--
				}
				pointer++
			}
			if exists {
				res = append(res, i)
			}
		}
	}

	return res
}
