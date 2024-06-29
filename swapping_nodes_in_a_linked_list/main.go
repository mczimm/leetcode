package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	head := &ListNode{Val: 1, Next: two}
	fmt.Println(swapNodes(head, 2)) //[1,4,3,2,5]
}

func swapNodes(head *ListNode, k int) *ListNode {
	var headS []int

	for head != nil {
		headS = append(headS, head.Val)
		head = head.Next
	}

	for i := 0; i < len(headS); i++ {
		if k-1 == i {
			tmp := headS[i]
			headS[i] = headS[len(headS)-k]
			headS[len(headS)-k] = tmp
		}
	}
	return sliceToLinkedList(headS)
}

func sliceToLinkedList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{Val: s[0], Next: nil}
	curr := head
	for _, v := range s[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}
