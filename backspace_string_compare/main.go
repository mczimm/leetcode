package main

import "fmt"

func main() {
	s := "a##c"
	t := "#a#c"
	fmt.Println(backspaceCompare(s, t)) //true
	s = "ab#c"
	t = "ad#c"
	fmt.Println(backspaceCompare(s, t)) //true
	s = "ab##"
	t = "c#d#"
	fmt.Println(backspaceCompare(s, t)) //true
	s = "a#c"
	t = "b"
	fmt.Println(backspaceCompare(s, t)) //false
}

func backspaceCompare(s string, t string) bool {
	currIndS := len(s) - 1
	currIndT := len(t) - 1
	for currIndS >= 0 || currIndT >= 0 {
		currIndS = getCurrentIndex(s, currIndS)
		currIndT = getCurrentIndex(t, currIndT)

		if currIndS < 0 && currIndT < 0 {
			return true
		}
		if currIndS < 0 || currIndT < 0 {
			return false
		}
		if s[currIndS] != t[currIndT] {
			return false
		}
		currIndS--
		currIndT--
	}

	return true
}

func getCurrentIndex(s string, i int) int {
	skip := 0

	for i >= 0 {
		if s[i] == byte('#') {
			skip++
		} else if skip > 0 {
			skip--
		} else {
			break
		}
		i--
	}
	return i
}
