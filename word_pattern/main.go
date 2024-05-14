package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  //true
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  //false
	fmt.Println(wordPattern("abba", "dog cat cat fish")) //false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  //false
}

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")

	if len(pattern) != len(ss) {
		return false
	}

	sm := make(map[string]byte)
	pm := make(map[byte]string)
	for i := 0; i < len(ss); i++ {
		if _, ok := pm[pattern[i]]; !ok {
			if _, ok = sm[ss[i]]; ok {
				return false
			} else {
				sm[ss[i]] = pattern[i]
				pm[pattern[i]] = ss[i]
			}
		} else {
			if pm[pattern[i]] != ss[i] {
				return false
			}
		}
	}
	return true
}
