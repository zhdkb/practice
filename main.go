package main

import (
	"fmt"
	"test/list"
)

func main() {
	list1 := &list.ListNode{Val: 1}
	list1.Next = &list.ListNode{Val: 2}
	list1.Next.Next = &list.ListNode{Val: 3}
	list1.Next.Next.Next = &list.ListNode{Val: 4}
	//list1.Next.Next.Next.Next = &list.ListNode{Val: 5}
	list1.Next.Next.Next.Next = list1
	// list1.Next.Next.Next = list1


	result := list.ChangePretoNext(list1)

	fmt.Println(result.Val)
	result = result.Next
	for result != list1 {
		fmt.Println(result.Val)
		result = result.Next
	}
}
