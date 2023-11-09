package main

import "fmt"

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd')) //dcbaefd
	fmt.Println(reversePrefix("xyxzxe", 'z'))  //zxyxxe
	fmt.Println(reversePrefix("abcd", 'z'))    //abcd
}

func reversePrefix(word string, ch byte) string {
	var res []byte
	var left, right int

	for right < len(word) {
		if word[right] == ch {
			res = reverseSubstring2(word[left : right+1])
			res = append(res, word[right+1:]...)
			return string(res)
		}
		right++
	}
	return word
}

func reverseSubstring(sub string) []byte {
	var res []byte
	for i := len(sub) - 1; i >= 0; i-- {
		res = append(res, sub[i])
	}
	return res
}

func reverseSubstring2(sub string) []byte {
	res := make([]byte, len(sub))
	var idx int
	for i := len(sub) - 1; i >= 0; i-- {
		res[idx] = sub[i]
		idx++
	}
	return res
}
