package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	three1 := &ListNode{Val: 4}
	two1 := &ListNode{Val: 3, Next: three1}
	head1 := &ListNode{Val: 1, Next: two1}

	three2 := &ListNode{Val: 4}
	two2 := &ListNode{Val: 2, Next: three2}
	head2 := &ListNode{Val: 1, Next: two2}

	res := mergeTwoLists(head1, head2)

	fmt.Println(res)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{Val: 0}
	result := merged

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			merged.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			merged.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		merged = merged.Next
	}

	//the list1 and list2 two can be different length
	//that's why we need to check it one more time separately
	for list1 != nil {
		merged.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		merged = merged.Next
	}

	for list2 != nil {
		merged.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		merged = merged.Next
	}

	return result.Next
}
