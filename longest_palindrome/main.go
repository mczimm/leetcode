package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd")) //7
	fmt.Println(longestPalindrome("a"))        //1
	fmt.Println(longestPalindrome("bb"))       //2
	fmt.Println(longestPalindrome("ccc"))      //3
}

// My first attempt
func longestPalindrome(s string) int {
	d := make(map[byte]int)
	var res int

	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}
	for k, v := range d {
		curr := v - (v % 2)
		res += curr
		d[k] -= curr
		if d[k] == 0 {
			delete(d, k)
		}
	}
	if len(d) <= 0 {
		return res
	} else {
		return res + 1
	}
}
