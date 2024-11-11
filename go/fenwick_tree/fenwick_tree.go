package fenwick_tree

type FenwickTree struct {
	tree []int
	n    int
}

func build(tree *FenwickTree, n int) {
	tree.n = n
	tree.tree = make([]int, n+1)
}

func (tree *FenwickTree) update(index, x int) {
	for i := index; i <= tree.n; i += (i & (-i)) {
		tree.tree[i] += x
	}
}

func (tree *FenwickTree) query(index int) int {
	res := 0
	for i := index; i > 0; i -= (i & (-i)) {
		res += tree.tree[i]
	}
	return res
}
