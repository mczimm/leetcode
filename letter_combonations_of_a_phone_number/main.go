package main

import "fmt"

func main() {
	fmt.Printf("%s", letterCombinations("23")) //["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Printf("%s", letterCombinations(""))   //[]
	fmt.Printf("%s", letterCombinations("2"))  //["a", "b", "c"]
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dm := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var res []string
	var backtrack func(int, string)

	backtrack = func(index int, curr string) {
		if len(curr) == len(digits) {
			res = append(res, curr)
			return
		}
		for _, letter := range dm[digits[index]] {
			println(string(letter))
			backtrack(index+1, curr+string(letter))
		}
	}

	backtrack(0, "")
	return res
}
