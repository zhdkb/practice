package main

import (
	"fmt"
	"test/list"
)

func main() {
	list1 := &list.ListNode{Val: 5}
	list1.Next = &list.ListNode{Val: 3}
	list1.Next.Next = &list.ListNode{Val: 2}
	list1.Next.Next.Next = &list.ListNode{Val: 4}
	list1.Next.Next.Next.Next = &list.ListNode{Val: 1}
	result := list.DoMergeSort(list1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
