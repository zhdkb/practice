package tree

type Nodelist struct {
	Val         int
	Left, Right *Nodelist
}

func Preorder(root *Nodelist) []int {
	result := make([]int, 0)
	store := []*Nodelist{root}

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

func Inorder(root *Nodelist) []int {
	result := make([]int, 0)
	store := make([]*Nodelist, 0)

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

func Postorder(root *Nodelist) []int {
	result := make([]int, 0)
	store := []*Nodelist{root}
	var pre *Nodelist

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
