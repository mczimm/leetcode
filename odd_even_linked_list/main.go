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
	fmt.Println(oddEvenList(head)) //[1,3,5,2,4]
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	enevHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = enevHead
	return head
}
