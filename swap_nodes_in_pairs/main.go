package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	four := &ListNode{Val: 4}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	head := &ListNode{Val: 1, Next: two}
	println(swapPairs(head)) //[2,1,4,3]
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		prev.Next = next
		cur.Next = next.Next
		next.Next = cur

		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}
