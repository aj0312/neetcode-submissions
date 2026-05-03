/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	list3 := &ListNode{}
	temp := list3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list3.Next = list1
			list3 = list3.Next
			list1 = list1.Next
		} else {
			list3.Next = list2
			list3 = list3.Next
			list2 = list2.Next
		}
	}
	if list2 == nil {
		list3.Next = list1
	}
	if list1 == nil {
		list3.Next = list2
	}
	return temp.Next
}
