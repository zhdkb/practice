package main

import (
	"fmt"
	"test/tree"
)

func main() {
	root := &tree.TreeNode{Val: 1}
	root.Left = &tree.TreeNode{Val: 2}
	root.Right = &tree.TreeNode{Val: 3}
	root.Left.Left = &tree.TreeNode{Val: 4}
	root.Left.Right = &tree.TreeNode{Val: 5}
	root.Right.Right = &tree.TreeNode{Val: 6}
	root.Left.Right.Left = &tree.TreeNode{Val: 7}
    root.Left.Right.Left.Left = &tree.TreeNode{Val: 8}
    root.Left.Right.Left.Left.Left = &tree.TreeNode{Val: 9}
    root.Left.Right.Left.Left.Left.Left = &tree.TreeNode{Val: 10}

    result := tree.GetLeafNodes(root)
    fmt.Println(result)
}
