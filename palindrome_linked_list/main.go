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
	fmt.Println(isPalindrome(head))
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	middle := middleNode(head)
	reverse := reverseList(middle)

	for reverse != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
