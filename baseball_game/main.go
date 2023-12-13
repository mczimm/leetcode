package main

import (
	"fmt"
	"strconv"
)

func main() {
	//30
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(operations []string) int {
	var tmp []int
	var res int
	for _, v := range operations {
		switch rune(v[0]) {
		case 'C':
			tmp = tmp[:len(tmp)-1]
		case 'D':
			tmp = append(tmp, tmp[len(tmp)-1]*2)
		case '+':
			tmp = append(tmp, tmp[len(tmp)-1]+tmp[len(tmp)-2])
		default:
			t, _ := strconv.Atoi(v)
			tmp = append(tmp, t)
		}
	}
	for i := 0; i < len(tmp); i++ {
		res += tmp[i]
	}
	return res
}
