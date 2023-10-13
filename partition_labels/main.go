package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij" //[9,7,8]
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	var res []int
	tmp := make(map[rune]int)
	part_end, part_len := 0, 0

	for i, v := range s {
		tmp[v] = i
	}

	for i, v := range s {
		part_end = Max(part_end, tmp[v])
		part_len += 1

		if part_end == i {
			res = append(res, part_len)
			part_len = 0
		}
	}

	return res
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
