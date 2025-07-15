package tree

/*
** 将二叉树的叶子结点添加到线性表中
*/
func GetLeafNodes(root *TreeNode) []int {
	result := make([]int, 0)
	store := []*TreeNode{root}
	for len(store) != 0 {
		node := store[len(store) - 1]
		store = store[: len(store) - 1]
		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
		}
		if node.Right != nil {
			store = append(store, node.Right)
		}
		if node.Left != nil {
			store = append(store, node.Left)
		}
	}

	return result
}
