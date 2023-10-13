package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	four := &ListNode{Val: 1}
	three := &ListNode{Val: 2, Next: four}
	two := &ListNode{Val: 2, Next: three}
	head := &ListNode{Val: 1, Next: two}
	fmt.Println(removeElements(head, 1))
}

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{Val: 0}
	result.Next = head
	current := result

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return result.Next
}
