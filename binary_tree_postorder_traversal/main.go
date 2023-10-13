package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int

	travers(root, &ans)
	return ans
}

func travers(node *TreeNode, ans *[]int) {
	if node != nil {
		travers(node.Left, ans)
		travers(node.Right, ans)
		*ans = append(*ans, node.Val)
	}
}
