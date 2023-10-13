package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("aa", "ab"))   //false
	fmt.Println(canConstruct("aa", "aab"))  //true
	fmt.Println(canConstruct("aab", "baa")) //true
}

func canConstruct(ransomNote string, magazine string) bool {
	//JUST FOR EXAMPLE HOW TO SORT
	//magazineSortedRune := []rune(magazine)
	//sort.Slice(magazineSortedRune, func(i, j int) bool {
	//	return magazineSortedRune[i] < magazineSortedRune[j]
	//})
	magazineMap := make(map[rune]int)

	for _, v := range magazine {
		if _, ok := magazineMap[v]; !ok {
			magazineMap[v] = 1
		} else {
			magazineMap[v]++
		}
	}

	for _, v := range ransomNote {
		if _, ok := magazineMap[v]; !ok {
			return false
		} else {
			magazineMap[v]--
			if magazineMap[v] <= 0 {
				delete(magazineMap, v)
			}
		}
	}
	return true
}
