package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(countSegments("Hello, my name is John"))   //5
	fmt.Println(countSegments("Hello"))                    //1
	fmt.Println(countSegments("                "))         //0
	fmt.Println(countSegments(""))                         //0
	fmt.Println(countSegments(", , , ,        a, eaefa"))  //6
	fmt.Println(countSegments(",,,,,ae%333**** 7 8 7 10")) //5
}

func countSegments(s string) int {
	if len(s) <= 0 {
		return 0
	}
	tmp := strings.Split(s, " ")
	var cnt int
	for _, v := range tmp {
		if len(v) > 0 && (unicode.IsLetter(rune(v[0])) || unicode.IsDigit(rune(v[0])) || strings.ContainsAny(v, "!@#$%^&*()_+-=',.:")) {
			cnt++
		}
	}
	return cnt
}
