package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	if !strings.Contains(haystack, needle) {
		return -1
	} else {
		for i := 0; i < len(haystack); i++ {

			if haystack[i] == needle[0] {
				tmp := make([]byte, len(needle))
				tmp = []byte(haystack[i : i+len(needle)])
				if string(tmp) == needle {
					return i
				}
			}
		}
	}

	return -1
}

func strStr2(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	var found bool

	for w := 0; w <= h-n; w++ {
		for i := 0; i < n; i++ {
			if needle[i] != haystack[w+i] {
				found = false
				break
			} else {
				found = true
			}
		}
		if found {
			return w
		}
	}
	return -1
}
