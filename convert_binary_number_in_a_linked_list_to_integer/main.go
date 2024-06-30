package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	three := &ListNode{Val: 1}
	two := &ListNode{Val: 0, Next: three}
	head := &ListNode{Val: 1, Next: two}
	fmt.Println(getDecimalValue(head)) //5
}

func getDecimalValue(head *ListNode) int {
	var bin []int
	for head != nil {
		bin = append(bin, head.Val)
		head = head.Next
	}
	res := conv(bin)
	return res
}

func conv(binary []int) int {
	decimal := 0
	n := len(binary)
	for i, v := range binary {
		if v == 1 {
			decimal += int(math.Pow(2, float64(n-i-1)))
		}
	}

	return decimal
}
