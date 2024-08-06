package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(isValid("()")) //true
	fmt.Println(isValid("(]")) //false
	fmt.Println(isValid("{"))  //false
}

type stack struct {
	elem []rune
}

func (s *stack) Push(ss rune) {
	s.elem = append(s.elem, ss)
}

func (s *stack) Pop() {
	if len(s.elem) == 0 {
		return
	}
	s.elem = s.elem[:len(s.elem)-1]
}

func (s *stack) Top() (rune, error) {
	if len(s.elem) == 0 {
		return -1, fmt.Errorf("empty")
	}
	return s.elem[len(s.elem)-1], nil
}

func isValid(s string) bool {
	var st stack
	matchBracket := make(map[rune]rune)
	matchBracket[')'] = '('
	matchBracket[']'] = '['
	matchBracket['}'] = '{'

	for _, c := range s {
		if openBracket, ok := matchBracket[c]; ok {
			lastBracket, err := st.Top()
			if err != nil {
				return false
			}
			if openBracket != lastBracket {
				return false
			}
			st.Pop()
		} else {
			st.Push(c)
		}
	}
	return len(st.elem) == 0
}

func isValid2(s string) bool {
	if len(s) <= 1 || len(s) >= int(math.Pow10(4)) {
		return false
	}

	if strings.ContainsAny(s[0:1], ")]}") {
		return false
	}

	var res []string

	if strings.ContainsAny(s, "()[]{}") {
		for _, v := range s {
			switch {
			case strings.ContainsAny(string(v), "([{"):
				res = append(res, string(v))
			case string(v) == ")" && len(res) > 0 && res[len(res)-1] == "(":
				res = res[:len(res)-1]
			case string(v) == "}" && len(res) > 0 && res[len(res)-1] == "{":
				res = res[:len(res)-1]
			case string(v) == "]" && len(res) > 0 && res[len(res)-1] == "[":
				res = res[:len(res)-1]
			default:
				return false
			}
		}
	}

	return len(res) == 0
}
