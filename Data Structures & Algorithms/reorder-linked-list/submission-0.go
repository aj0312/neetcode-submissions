/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
		return
	}
	node := head
	nodes := []*ListNode{}
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	for len(nodes) > 2  {
		firstNode := nodes[0]
		secondLastNode := nodes[len(nodes)-2]
		lastNode := nodes[len(nodes)-1]
		secondLastNode.Next = lastNode.Next
		lastNode.Next = firstNode.Next
		firstNode.Next = lastNode
		nodes = nodes[1:len(nodes)-1]
	}

}
