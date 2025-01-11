package main

func main() {
	println(canConstruct2("annabelle", 2)) // true
	println(canConstruct2("leetcode", 3))  // false
	println(canConstruct2("cr", 7))        // false
	println(canConstruct2("cr", 7))        // false
}

func canConstruct(s string, k int) bool {

	if len(s) < k {
		return false
	}
	if len(s) == k {
		return true
	}
	var freq [26]int
	var oddCount int
	for _, v := range s {
		freq[v-'a']++
	}
	for _, v := range freq {
		if v%2 == 1 {
			oddCount++
		}
	}
	println("fuck")
	return oddCount <= k
}

func isPalindrome(s string) bool {
	right := len(s) - 1

	for left := 0; left < len(s); left++ {
		if s[left] != s[right] {
			return false
		}
		right--
	}
	return true
}

func canConstruct2(s string, k int) bool {
	if len(s) < k {
		return false
	}
	if len(s) == k {
		return true
	}
	freq := make(map[int32]int)
	var cnt int

	for _, v := range s {
		freq[v]++
	}
	for _, v := range freq {
		if v%2 == 1 {
			cnt++
		}
	}
	return cnt <= k
}
