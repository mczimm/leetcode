package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(bitwiseComplement(5)) //2
}

func bitwiseComplement(n int) int {
	var res []byte
	b1 := strconv.FormatInt(int64(n), 2)
	for _, v := range b1 {
		if v == '1' {
			res = append(res, '0')
		} else {
			res = append(res, '1')
		}
	}
	ress, _ := strconv.ParseInt(string(res), 2, 32)
	return int(ress)
}
