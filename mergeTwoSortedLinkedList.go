package main

/**
* [LeetCode 21]
**/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	temp := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1.Next
			temp = temp.Next
		} else {
			temp.Next = list2.Next
			temp = temp.Next
		}

	}

	for list1 != nil {
		temp.Next = list1
		list1 = list1.Next
		temp = temp.Next
	}

	for list2 != nil {
		temp.Next = list2
		list2 = list2.Next
		temp = temp.Next
	}

	res = res.Next
	return res
}
