package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog")) // true
}

func checkIfPangram(sentence string) bool {
	const engLength int = 26
	res := make(map[rune]struct{})
	for _, v := range sentence {
		res[v] = struct{}{}
	}
	if len(res) == engLength {
		return true
	}
	return false
}
