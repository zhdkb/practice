package list

func Calrepeat(l1, l2 *ListNode) *ListNode {
	l := &ListNode{Val: 1}
	store := l
	result := l1
	judge := false

	for l1 != nil && l2 != nil {
		if l1.Val == l2.Val {
			if l1.Val == l.Val && judge {
				l1 = l1.Next
				l2 = l2.Next
				continue
			}
			l.Next = &ListNode{Val: l1.Val}
			l = l.Next
			judge = true
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Val > l2.Val {
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
	}

	result = store.Next
	return result
}
