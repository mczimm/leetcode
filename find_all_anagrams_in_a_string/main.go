package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("baa", "aa"))         //[1]
	fmt.Println(findAnagrams("bab", "abc"))        //[]
}

// My attempt
func findAnagrams(s string, p string) []int {
	var res []int
	mapP := make(map[byte]int)
	mapS := make(map[byte]int)

	if len(s) < len(p) {
		return res
	}

	for i := 0; i < len(p); i++ {
		mapP[p[i]]++
	}

	for i := 0; i < len(s); i++ {
		if i-len(p) >= 0 {
			mapS[s[i-len(p)]]--

			if mapS[s[i-len(p)]] == 0 {
				delete(mapS, s[i-len(p)])
			}
		}

		mapS[s[i]]++

		if len(mapP) == len(mapS) {
			exists := true

			for k, val := range mapP {
				if v, ok := mapS[k]; !ok {
					exists = false
					break
				} else {
					if v != val {
						exists = false
						break
					}
				}
			}
			if exists {
				res = append(res, i+1-len(p))
			}
		}
	}

	return res
}
