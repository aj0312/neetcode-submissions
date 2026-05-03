/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodes := []*ListNode{}

	node := head

	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	prevIdx := len(nodes)-n-1
	if prevIdx < 0 {
		temp := head.Next
		head.Next = nil
		head = temp
		return head
	}

	prevNode := nodes[len(nodes)-n-1]
	temp := prevNode.Next
	prevNode.Next = temp.Next
	temp.Next = nil

	return head
}
