package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4, Next: three}
	first := &ListNode{Val: 2, Next: four}

	four2 := &ListNode{Val: 4}
	six := &ListNode{Val: 6, Next: four2}
	second := &ListNode{Val: 5, Next: six}
	println(addTwoNumbers(first, second))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		println(sum / 10)
		println(sum % 10)
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}
