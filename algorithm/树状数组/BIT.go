type BIT struct {
    tree []int 
    n int 
}

func build(tree *BIT, n int) {
    tree.n = n
    tree.tree = make([]int, n+1)
}

func (tree *BIT) update(index, x int) {
    for i:=index; i<=tree.n; i+=(i&(-i)) {
        tree.tree[i] += x 
    }
}

func (tree *BIT) query(index int) int {
    res := 0 
    for i:=index; i>0; i-=(i&(-i)) {
        res += tree.tree[i]
    }
    return res 
}