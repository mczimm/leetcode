/*
Given the root of a binary tree, return the average value of the nodes on each level in the form of an array.
Answers within 10-5 of the actual answer will be accepted.
*/

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	twol := &TreeNode{15, nil, nil}
	twor := &TreeNode{7, nil, nil}
	onel := &TreeNode{9, nil, nil}
	oner := &TreeNode{20, twol, twor}
	head := &TreeNode{3, onel, oner}

	fmt.Println(averageOfLevels(head))
}

// not working
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		res = append(res, float64(node.Val))
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
