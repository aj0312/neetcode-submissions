/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return true
	}
	nums := inOrderTraversal(root, []int{})

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}

func inOrderTraversal(root *TreeNode, numSlice []int) []int {
	if root == nil {
		return numSlice
	}
	numSlice = inOrderTraversal(root.Left, numSlice)
	numSlice = append(numSlice, root.Val)
	numSlice = inOrderTraversal(root.Right, numSlice)
	return numSlice
}
