package main

import "fmt"

func main() {
	//fmt.Println(longestPalindrome2("babad")) //bab
	//fmt.Println(longestPalindrome2("bab"))   //bab
	//fmt.Println(longestPalindrome2("bb"))          //bb
	fmt.Println(longestPalindrome2("aacabdkacaa")) //aca
	//fmt.Println(longestPalindrome2("cbbd"))  //bb
	//fmt.Println(longestPalindrome2("c"))     //c
	//fmt.Println(longestPalindrome2("ac"))    //a
	//fmt.Println(longestPalindrome2("acc"))   //cc
}

// brute-force
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	palidrome := func(i, j int) bool {
		left := i
		right := j - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if palidrome(j, i+j) {
				return s[j : i+j]
			}
		}
	}
	return ""
}

func longestPalindrome2(s string) string {
	var res string
	if len(s) == 1 {
		return s
	}

	for i := 0; i <= len(s); i++ { //"aacabdkacaa"
		l, r := i, i
		for l <= r {
			if r >= len(s) || l <= -1 {
				break
			}
			if s[l] != s[r] {
				break
			}
			if r-l+1 >= len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
		l, r = i, i+1
		for l <= r {
			if r >= len(s) || l <= -1 {
				break
			}
			if s[l] != s[r] {
				break
			}
			if r-l+1 > len(res) {
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}
