package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest")) //"s'teL ekat edoCteeL tsetnoc"
}

func reverseWords(s string) string {
	strSlice := strings.Split(s, " ")
	var res string

	for i, word := range strSlice {
		//tmp := word
		for j := len(word) - 1; j >= 0; j-- {
			//res += strings.Join([]string{string(tmp[j])}, "")
			res += string(word[j])
		}
		if i < len(strSlice)-1 {
			res += " "
		}
	}
	return res
}
