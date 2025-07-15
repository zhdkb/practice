package list

/*
** 删除单链表中的多余结点
*/
func Deletesame(list *ListNode) *ListNode {
	result := list
	var pre *ListNode

	for list != nil {
		node := result
		for pre != nil && node != list {
			if node.Val == list.Val {
				pre.Next = list.Next
				list = pre
				break
			}
			node = node.Next
		}
		pre = list
		list = list.Next
	}

	return result
}
