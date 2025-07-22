package list

func mergetwo(list1, list2 *ListNode) *ListNode {
	list := &ListNode{Val: 1}
	result := list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list.Next = list1
			list1 = list1.Next
		} else {
			list.Next = list2
			list2 = list2.Next
		}
		list = list.Next
	}

	if list1 != nil {
		list.Next = list1
	} else {
		list.Next = list2
	}

	return result.Next
}

func mergesort(first, end *ListNode) *ListNode {
	if first == end || first.Next == end {
		first.Next = nil
		return first
	}

	list1, list2 := first, first
	for list2 != end {
		list1 = list1.Next
		list2 = list2.Next
		if list2 != end {
			list2 = list2.Next
		}
	}

	return mergetwo(mergesort(first, list1), mergesort(list1, end))
}

func DoMergeSort(list *ListNode) *ListNode {
	return mergesort(list, nil)
}
