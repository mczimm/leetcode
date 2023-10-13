package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))  //true
	fmt.Println(isAnagram("rat", "cat"))          //false
	fmt.Println(isAnagram("ab", "a"))             //false
	fmt.Println(isAnagram("aacc", "ccac"))        //false
	fmt.Println(isAnagram2("anagram", "nagaram")) //true
	fmt.Println(isAnagram2("rat", "cat"))         //false
	fmt.Println(isAnagram2("ab", "a"))            //false
	fmt.Println(isAnagram2("aacc", "ccac"))       //false
}

// My approach
func isAnagram(s string, t string) bool {
	var sliceS []string
	var sliceT []string
	if len(t) != len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		sliceS = append(sliceS, string(s[i]))
		sliceT = append(sliceT, string(t[i]))
	}
	sort.Strings(sliceS)
	sort.Strings(sliceT)

	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] != sliceT[i] {
			return false
		}
	}
	return true
}

//From the solutions: https://leetcode.com/problems/valid-anagram/solutions/3501509/isanagram-go-array-map/?source=submission-ac

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp := make(map[rune]int)

	for _, v := range s {
		mp[v]++

	}
	for _, v := range t {
		mp[v]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
