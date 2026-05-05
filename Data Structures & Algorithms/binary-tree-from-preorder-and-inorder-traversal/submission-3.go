/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderMap := map[int]int{}

	for i, num := range inorder {
		inorderMap[num] = i
	}
	root, _ := buildBinaryTree(preorder, 0, len(inorder)-1, inorderMap)
	return root
}

func buildBinaryTree(preorder []int, start, end int, inorderMap map[int]int) (*TreeNode, []int) {
	if start > end {
		return nil, preorder
	}
	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	inorderIdx := inorderMap[root.Val]
	root.Left, preorder = buildBinaryTree(preorder, start, inorderIdx-1, inorderMap)
	root.Right, preorder = buildBinaryTree(preorder, inorderIdx+1, end, inorderMap)
	return root, preorder
}
