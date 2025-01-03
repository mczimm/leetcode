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
	println(swapPairs(head))  //[2,1,4,3]
	println(swapPairs2(head)) //[2,1,4,3]
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

//recursion
func swapPairs2(head *ListNode) *ListNode {
	// If the list has no node or has only one node left.
	if head == nil || head.Next == nil {
		return head
	}
	// Nodes to be swapped
	firstNode := head
	secondNode := head.Next
	// Swapping
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode
	// Now the head is the second node
	return secondNode
}
