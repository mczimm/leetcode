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
	fmt.Println(removeNthFromEnd(head, 2)) //[1,2,3,5]
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenListNode := 0
	result := &ListNode{}
	result.Next = head
	curr := result

	for head != nil {
		head = head.Next
		lenListNode++
	}

	if lenListNode == 1 {
		return head
	}

	for lenListNode-n > 0 {
		curr = curr.Next
		lenListNode--
	}
	curr.Next = curr.Next.Next
	return result.Next
}
