package list

/*
** 将双链表中的结点，根据Frep（访问数）进行排序
*/
func Visited(list *ListNode, num int) *ListNode {
	result := list
	for list != nil {
		if list.Val != num {
			list = list.Next
			continue
		}

		list.Freq++
		if list.Pre == nil {
			list = list.Next
			continue
		}
		n := list.Next
		for list.Pre != nil && list.Freq > list.Pre.Freq {
			// 交换相邻结点
			node := list.Pre
			list.Pre = node.Pre
			if node.Pre != nil {
				node.Pre.Next = list
			}
			node.Next = list.Next
			if list.Next != nil {
				list.Next.Pre = node
			}
			list.Next = node
			node.Pre = list
		}
		if list.Pre == nil {
			result = list
		}
		list = n
	}

	return result
}
