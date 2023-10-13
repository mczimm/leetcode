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

	for i := 0; i < len(strSlice); i++ {
		tmp := strSlice[i]
		for j := len(tmp) - 1; j >= 0; j-- {
			res += strings.Join([]string{string(tmp[j])}, "")
		}
		if i < len(strSlice)-1 {
			res += " "
		}
	}
	return res
}
