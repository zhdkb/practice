package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Preorder(root *TreeNode) []int {
	result := make([]int, 0)
	store := []*TreeNode{root}

	for len(store) != 0 {
		node := store[len(store) - 1]
		store = store[: len(store) - 1]
		result = append(result, node.Val)
		if node.Right != nil {
			store = append(store, node.Right)
		}
		if node.Left != nil {
			store = append(store, node.Left)
		}
	}

	return result
}

func Inorder(root *TreeNode) []int {
	result := make([]int, 0)
	store := make([]*TreeNode, 0)

	for root != nil || len(store) != 0 {
		for root != nil {
			store = append(store, root)
			root = root.Left
		}

		node := store[len(store) - 1]
		store = store[: len(store) - 1]
		result = append(result, node.Val)
		root = node.Right
	}

	return result
}

func Postorder(root *TreeNode) []int {
	result := make([]int, 0)
	store := []*TreeNode{root}
	var pre *TreeNode

	for len(store) != 0 {
		node := store[len(store) - 1]
		if (node.Left == nil && node.Right == nil) || ( pre != nil && (pre == node.Left || pre == node.Right)) {
			result = append(result, node.Val)
			store = store[: len(store) - 1]
			pre = node
		} else {
			if node.Right != nil {
				store = append(store, node.Right)
			}
			if node.Left != nil {
				store = append(store, node.Left)
			}
		}
	}

	return result
}
