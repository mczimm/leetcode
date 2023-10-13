/*
https://coderun.yandex.ru/problem/pin?currentPage=1&pageSize=10&rowNumber=1
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	var cnt int

	_, err := fmt.Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}

	var nails = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		_, err := fmt.Scan(&nails[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	//cnt = 6
	//1.nails = []int{3, 13, 12, 4, 14, 6}

	fmt.Println(threadLength(cnt, nails))
}

func threadLength(cnt int, nails []int) int {
	var res = make([]int, cnt)

	nailsSorted := bubbleSort(nails)
	res[1] = nailsSorted[1] - nailsSorted[0]

	if cnt > 2 {
		res[2] = nailsSorted[2] - nailsSorted[0]
	}

	for i := 3; i < cnt; i++ {
		res[i] = min(res[i-2], res[i-1]) + nailsSorted[i] - nailsSorted[i-1]
	}

	return res[cnt-1]
}

func bubbleSort(nails []int) []int {
	for i := 0; i < len(nails)-1; i++ {
		for j := 0; j < len(nails)-i-1; j++ {
			if nails[j] > nails[j+1] {
				nails[j], nails[j+1] = nails[j+1], nails[j]
			}
		}
	}
	return nails
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
