/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	nums := inOrderTraverse(root, []int{})
	return nums[k-1]
}

func inOrderTraverse(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = inOrderTraverse(root.Left, nums)
	nums = append(nums, root.Val)
	nums = inOrderTraverse(root.Right, nums)
	return nums
}