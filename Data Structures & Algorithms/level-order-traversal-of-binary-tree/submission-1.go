/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
	output := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []int{}
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		output = append(output, level)
	}
	return output
}
