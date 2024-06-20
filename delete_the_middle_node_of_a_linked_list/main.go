package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	seven := &ListNode{Val: 6}
	six := &ListNode{Val: 2, Next: seven}
	five := &ListNode{Val: 1, Next: six}
	four := &ListNode{Val: 7, Next: five}
	three := &ListNode{Val: 4, Next: four}
	two := &ListNode{Val: 3, Next: three}
	head := &ListNode{Val: 1, Next: two}
	fmt.Println(deleteMiddle(head)) //[1,3,4,1,2,6]
}

func deleteMiddle(head *ListNode) *ListNode {
	mid := findMiddle(head)
	curr := head

	dummy := &ListNode{0, head}
	prev := dummy

	for curr != nil {
		nextNode := curr.Next
		if curr == mid {
			prev.Next = nextNode
			break
		}
		prev = curr
		curr = curr.Next
	}
	return dummy.Next
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
