package main

import (
	"fmt"
	"test/tree"
)

func main() {
	root := &tree.Nodelist{Val: 1}
    root.Left = &tree.Nodelist{Val: 2}
    root.Right = &tree.Nodelist{Val: 3}
    root.Left.Left = &tree.Nodelist{Val: 4}
    root.Left.Right = &tree.Nodelist{Val: 5}
    root.Right.Left = &tree.Nodelist{Val: 6}
    root.Right.Right = &tree.Nodelist{Val: 7}
    root.Left.Left.Left = &tree.Nodelist{Val: 8}
    root.Left.Left.Right = &tree.Nodelist{Val: 9}
    root.Right.Right.Left = &tree.Nodelist{Val: 10}
    root.Right.Right.Right = &tree.Nodelist{Val: 11}

	pre := tree.Preorder(root)
	in := tree.Inorder(root)
	post := tree.Postorder(root)

	fmt.Println(pre)
	fmt.Println(in)
	fmt.Println(post)
}
