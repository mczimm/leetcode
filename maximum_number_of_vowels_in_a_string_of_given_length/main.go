package main

import "fmt"

func main() {
	fmt.Println(maxVowels("abciiidef", 3))                                               //3
	fmt.Println(maxVowels("aeiou", 2))                                                   //2
	fmt.Println(maxVowels("leetcode", 3))                                                //2
	fmt.Println(maxVowels("weallloveyou", 7))                                            //4
	fmt.Println(maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33))                     //7
	fmt.Println(maxVowels("ovkrzjmplhcloojfjrccmrscwqqhaackfyorqpzekrzvpynqpmgkqd", 48)) //7
	fmt.Println()
	fmt.Println(maxVowels2("abciiidef", 3))                                               //3
	fmt.Println(maxVowels2("aeiou", 2))                                                   //2
	fmt.Println(maxVowels2("leetcode", 3))                                                //2
	fmt.Println(maxVowels2("weallloveyou", 7))                                            //4
	fmt.Println(maxVowels2("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33))                     //7
	fmt.Println(maxVowels2("ovkrzjmplhcloojfjrccmrscwqqhaackfyorqpzekrzvpynqpmgkqd", 48)) //7
}

// First attempt 10.11.23
func maxVowels(s string, k int) int {
	var res, vowels int
	//calculate all vowels before k
	for i := 0; i < k; i++ {
		if isVowel(rune(s[i])) {
			vowels++
		}
	}
	if k == len(s)-1 || vowels == k {
		return vowels
	}
	//current maximum of vowels
	res = vowels
	//continue from k
	for i := k; i < len(s); i++ {
		//before we slide window to the right we must insure if we don't have a vowel
		//otherwise delete it because it is not our window anymore
		if isVowel(rune(s[i-k])) {
			vowels--
		}
		//if last element in window is vowel
		if isVowel(rune(s[i])) {
			vowels++
		}
		if vowels > res {
			res = vowels
		}
	}
	return res
}

// From the solutions
func maxVowels2(s string, k int) int {
	var res, vowels int

	for right := 0; right < len(s); right++ {
		if isVowel(rune(s[right])) {
			vowels++
		}
		if right >= k && isVowel(rune(s[right-k])) { //first element of window
			vowels--
		}
		if right >= k-1 && vowels > res { //whole window (k-1 because array starts from 0)
			res = vowels
		}
		if res == k {
			return res
		}
	}
	if vowels > res && vowels <= k {
		return vowels
	}
	return res
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
