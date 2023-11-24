package main

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))                 // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))          // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))                  // 0
	fmt.Println(maxNumberOfBalloons("balllllllllllloooooooooon")) // 1
	fmt.Println(maxNumberOfBalloons("balon"))                     // 0
}

func maxNumberOfBalloons(text string) int {
	res := len(text)
	balloons := map[rune]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}
	for _, v := range text {
		if v == 'b' || v == 'a' || v == 'l' || v == 'o' || v == 'n' {
			balloons[v]++
		}
	}

	for i, v := range balloons {
		if i == 'l' || i == 'o' {
			v = v / 2
		}
		if v < res {
			res = v
		}
	}

	return res
}
