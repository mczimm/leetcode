package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	var target byte
	target = 'g'
	fmt.Println(string(nextGreatestLetter(letters, target)))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}

	return letters[0]
}
