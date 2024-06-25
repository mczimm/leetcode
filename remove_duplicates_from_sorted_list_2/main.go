package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	seven := &ListNode{Val: 5}
	six := &ListNode{Val: 4, Next: seven}
	five := &ListNode{Val: 4, Next: six}
	four := &ListNode{Val: 3, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	head := &ListNode{Val: 1, Next: two}
	fmt.Println(deleteDuplicates(head)) //[1,2,5]
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	result := &ListNode{0, head}
	prev := result

	for head != nil {
		// If it's a beginning of duplicates sublist
		// skip all duplicates
		if head.Next != nil && head.Val == head.Next.Val {
			// move till the end of duplicates sublist
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// Skip all duplicates
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return result.Next
}
