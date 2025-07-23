package main

import (
	"fmt"
	"test/list"
)

func main() {
	list1 := &list.ListNode{Val: 1}
	list1.Next = &list.ListNode{Val: 1}
	list1.Next.Next = &list.ListNode{Val: 3}
	list1.Next.Next.Next = &list.ListNode{Val: 4}
	list1.Next.Next.Next.Next = &list.ListNode{Val: 5}

	list2 := &list.ListNode{Val: 1}
	list2.Next = &list.ListNode{Val: 1}
	list2.Next.Next = &list.ListNode{Val: 4}
	list2.Next.Next.Next = &list.ListNode{Val: 7}
	list2.Next.Next.Next.Next = &list.ListNode{Val: 8}


	result := list.Calrepeat(list1, list2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
