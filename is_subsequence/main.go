package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))    //true
	fmt.Println(isSubsequence("axc", "ahbgdc"))    //false
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa")) //false
}

func isSubsequence(s string, t string) bool {
	cnt := 0
	ind := -1

	for i := 0; i < len(s); i++ { //0 1
		for ind < len(t)-1 { //0 1 2
			ind++
			if s[i] == t[ind] {
				cnt++
				break
			}
		}
	}
	if cnt == len(s) {
		return true
	}
	return false
}
