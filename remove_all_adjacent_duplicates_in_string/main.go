package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))    //ca
	fmt.Println(removeDuplicates("azxxzy"))    //ay
	fmt.Println(removeDuplicates("aaaaaaaa"))  // ""
	fmt.Println(removeDuplicates1("abbaca"))   //ca
	fmt.Println(removeDuplicates1("azxxzy"))   //ay
	fmt.Println(removeDuplicates1("aaaaaaaa")) // "
}

// My first approach, but it work very long (Time Limit Exceeded)
func removeDuplicates(s string) string {
	left, right := 0, 1
	sSlice := strings.Split(s, "")
	for right < len(sSlice) { // a b b a c a
		if sSlice[left] == sSlice[right] {
			sSlice = remove(sSlice, left)
			sSlice = remove(sSlice, right-1)
			left, right = 0, 1
		} else {
			left++
			right++
		}
	}
	return strings.Join(sSlice, "")
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// My second approach, works well
func removeDuplicates1(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		if len(res) > 0 && s[i] == res[len(res)-1] {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
