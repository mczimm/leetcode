package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("1a"))              //1
	fmt.Println(myAtoi("42"))              //42
	fmt.Println(myAtoi("   -42"))          //-42
	fmt.Println(myAtoi("4193 with words")) //4193
	fmt.Println(myAtoi("words and 987"))   //0
	fmt.Println(myAtoi("-91283472332"))    //-2147483648
	fmt.Println(myAtoi("3.14159"))         //3
	fmt.Println(myAtoi("+1"))              //1
	fmt.Println(myAtoi("+-12"))            //0
	fmt.Println(myAtoi("0-42a1234"))       //0
	fmt.Println(myAtoi("   +0 123"))       //0
	fmt.Println(myAtoi("-5-"))             //-5
	fmt.Println(myAtoi("  +  413"))        //0
	fmt.Println(myAtoi(" ++1"))            //0
}

// My first attempt
func myAtoi(s string) int {
	res := 0
	var tmp []string
	var negative bool
	var positive bool

	const (
		signed32   int = 2147483648 - 1
		unsigned32 int = -2147483648
	)

	for i, v := range s {
		if v == '-' {
			if i >= 1 && len(tmp) > 0 {
				break
			} else if negative {
				return 0
			} else {
				negative = true
			}
		}
		if v == '+' {
			if i >= 1 && len(tmp) > 0 {
				break
			} else if positive {
				return 0
			} else {
				positive = true
			}
		}
		if unicode.IsLetter(v) || v == '.' {
			break
		}
		if positive && negative {
			return 0
		}
		if unicode.IsDigit(v) {
			tmp = append(tmp, string(v))
		}
		if unicode.IsSpace(v) && len(tmp) > 0 {
			break
		}
		if unicode.IsSpace(v) && (positive || negative) {
			return 0
		}
	}

	res, _ = strconv.Atoi(strings.Join(tmp, ""))

	if negative {
		res = res * -1
	}

	if res > signed32 {
		return signed32
	} else if res < unsigned32 {
		return unsigned32
	} else {
		return res
	}
}
