package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))      //3
	fmt.Println(romanToInt("LVIII"))    //58
	fmt.Println(romanToInt("MCMXCIV"))  //1994
	fmt.Println(romanToInt2("III"))     //3
	fmt.Println(romanToInt2("LVIII"))   //58
	fmt.Println(romanToInt2("MCMXCIV")) //1994
}

// My first attempt
func romanToInt(s string) int {
	var res int

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "V" && i > 0 && string(s[i-1]) == "I" {
			res += 4
			i--
			continue
		} else if string(s[i]) == "X" && i > 0 && string(s[i-1]) == "I" {
			res += 9
			i--
			continue
		} else if string(s[i]) == "L" && i > 0 && string(s[i-1]) == "X" {
			res += 40
			i--
			continue
		} else if string(s[i]) == "C" && i > 0 && string(s[i-1]) == "X" {
			res += 90
			i--
			continue
		} else if string(s[i]) == "D" && i > 0 && string(s[i-1]) == "C" {
			res += 400
			i--
			continue
		} else if string(s[i]) == "M" && i > 0 && string(s[i-1]) == "C" {
			res += 900
			i--
			continue
		}

		if string(s[i]) == "I" {
			res += 1
			continue
		} else if string(s[i]) == "V" {
			res += 5
			continue
		} else if string(s[i]) == "X" {
			res += 10
			continue
		} else if string(s[i]) == "L" {
			res += 50
			continue
		} else if string(s[i]) == "C" {
			res += 100
			continue
		} else if string(s[i]) == "D" {
			res += 500
			continue
		} else if string(s[i]) == "M" {
			res += 1000
			continue
		}

	}

	return res
}

// From the solutions: https://leetcode.com/problems/roman-to-integer/solutions/4061855/0-ms-beats-100/
var symbol = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getSymbolValue(val uint8) int {
	return symbol[fmt.Sprintf("%c", val)]
}

func romanToInt2(str string) int {
	lastIndex := len(str) - 1
	value := getSymbolValue(str[lastIndex])
	for i := 0; i < lastIndex; i++ {
		int1 := getSymbolValue(str[i])
		int2 := getSymbolValue(str[i+1])
		if int1 >= int2 {
			value += int1
		} else {
			if int1*5 == int2 || int1*10 == int2 {
				value -= int1
			}
		}
	}
	return value
}
