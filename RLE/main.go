package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "AAAABBBCCXYZDDDDEEEFFFAAAAAAXX" // A4B3C2XYZD4E3F3A6X2
	fmt.Println(rle(str))
	fmt.Println(rle2(str))
}

func rle(s string) string {
	right := 1
	cnt := 1
	res := ""
	lastCh := ""

	for left := 0; left < len(s)-1; left++ {
		if s[left] == s[right] {
			cnt++
		} else {
			if cnt > 1 {
				res += string(s[left]) + fmt.Sprint(cnt)
				cnt = 1
			} else {
				res += string(s[left])
			}
		}
		right++
	}

	lastCh = s[len(s)-1:]

	if cnt > 1 {
		res += lastCh + fmt.Sprint(cnt)
	} else {
		res += lastCh
	}
	return res
}

func rle2(s string) string {
	if len(s) == 0 {
		fmt.Println("")
		return ""
	}

	// Validate input
	for _, ch := range s {
		if !unicode.IsUpper(ch) {
			fmt.Println("Invalid input: Only uppercase English letters are allowed.")
			return ""
		}
	}

	var builder strings.Builder
	cnt := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			builder.WriteByte(s[i-1])
			if cnt > 1 {
				builder.WriteString(fmt.Sprint(cnt))
			}
			cnt = 1
		}
	}

	// Add the last character and its count
	builder.WriteByte(s[len(s)-1])
	if cnt > 1 {
		builder.WriteString(fmt.Sprint(cnt))
	}

	return builder.String()
}
