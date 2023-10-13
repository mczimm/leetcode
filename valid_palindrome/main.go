/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) //true
	fmt.Println(isPalindrome("race a car"))                     //false
	fmt.Println(isPalindrome(" "))                              //true
	fmt.Println(isPalindrome("0P"))                             //false
}

func isPalindrome(s string) bool {
	converted := makeASCII(s)
	right := len(converted) - 1

	for left := 0; left < len(converted)-1; left++ {
		if converted[left] == converted[right] {
			right--
			continue
		} else {
			return false
		}
	}

	return true
}

func makeASCII(s string) string {
	var resASCII string
	for _, v := range s {
		if unicode.IsLetter(v) {
			resASCII += string(unicode.ToLower(v))
		} else if n, err := strconv.Atoi(string(v)); err == nil {
			resASCII += fmt.Sprint(n)
		} else {
			continue
		}

	}
	return resASCII
}
