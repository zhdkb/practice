package main

import (
	"fmt"
	"test/list"
)

func main() {
	l1 := &list.ListNode{
        Val: 1,
        Freq: 0,
    }

    l2 := &list.ListNode{
        Val: 2,
        Freq: 0,
    }

    l3 := &list.ListNode{
        Val: 3,
        Freq: 0,
    }

    l4 := &list.ListNode{
        Val: 4,
        Freq: 0,
    }
    l1.Next = l2
    l2.Pre = l1
    l2.Next = l3
    l3.Pre = l2
    l3.Next = l4
    l4.Pre = l3

    node := list.Visited(l1, 2)
    node = list.Visited(node, 3)
    node = list.Visited(node, 4)
    node = list.Visited(node, 4)
    //node = list.Visited(node, 3)
    for node != nil {
        fmt.Println(node.Val)
        node = node.Next
    }
}
