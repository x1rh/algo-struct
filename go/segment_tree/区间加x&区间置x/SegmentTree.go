package main

const (
	EMPTY = 0x3f3f3f3f3f3f // 一定要确保这个数用不到
)

// int通常是8 Byte，如果添加一个布尔类型的标记，那么内存对齐后，每个结构体将浪费7 Byte
// 在某些恶心的题目里为了节约内存，让addv和setv这样的标记字段包含两种语义，一是值，二是标记。这要求能区分出值与标记。
// 此外, go中int是64bit的
type Node struct {
	sumv int
	addv int
	setv int
	l, r int
}

type SegmentTree struct {
	tree []Node // slice自动扩容存在开销
}

// 0号节点是根节点，即tree[0]
func NewSegmentTree() *SegmentTree {
	st := &SegmentTree{}
	st.tree = append(st.tree, NewNode())
	return st
}

func NewNode() Node {
	return Node{
		sumv: 0,
		addv: EMPTY,
		setv: EMPTY,
		l:    0,
		r:    0,
	}
}

// pushUp 要求l和r一定存在，不检查l和r是否存在，因为pushUp配合pushDown使用
func (st *SegmentTree) pushUp(p int) {
	l := st.tree[p].l
	r := st.tree[p].r
	st.tree[p].sumv = st.tree[l].sumv + st.tree[r].sumv
}

func (st *SegmentTree) pushDown(p, Len int) {
	if st.tree[p].l == 0 { // 利用初始值0判断节点p还没有分配子节点
		st.tree[p].l = len(st.tree)
		st.tree = append(st.tree, NewNode())
	}
	if st.tree[p].r == 0 {
		st.tree[p].r = len(st.tree)
		st.tree = append(st.tree, NewNode())
	}

	l := st.tree[p].l
	r := st.tree[p].r

	if st.tree[p].setv != EMPTY {
		st.tree[l].setv = st.tree[p].setv
		st.tree[r].setv = st.tree[p].setv
		st.tree[l].addv = EMPTY
		st.tree[r].addv = EMPTY
		st.tree[l].sumv = st.tree[p].setv * (Len - (Len / 2))
		st.tree[r].sumv = st.tree[p].setv * (Len / 2)
		st.tree[p].setv = EMPTY
	}
	if st.tree[p].addv != EMPTY {
		if st.tree[l].addv != EMPTY {
			st.tree[l].addv += st.tree[p].addv
		} else {
			st.tree[l].addv = st.tree[p].addv
		}

		if st.tree[r].addv != EMPTY {
			st.tree[r].addv += st.tree[p].addv
		} else {
			st.tree[r].addv = st.tree[p].addv
		}

		st.tree[l].sumv += st.tree[p].addv * (Len - Len/2)
		st.tree[r].sumv += st.tree[p].addv * (Len / 2)
		st.tree[p].addv = EMPTY
	}
}

func (st *SegmentTree) Add(p, l, r, a, b, v int) {
	if a <= l && r <= b {
		st.tree[p].sumv += v * (r - l + 1)
		if st.tree[p].addv != EMPTY {
			st.tree[p].addv += v
		} else {
			st.tree[p].addv = v
		}
	} else {
		st.pushDown(p, r-l+1)
		m := (l + r) / 2
		if a <= m {
			st.Add(st.tree[p].l, l, m, a, b, v)
		}
		if m < b {
			st.Add(st.tree[p].r, m+1, r, a, b, v)
		}
		st.pushUp(p)
	}
}

func (st *SegmentTree) Set(p, l, r, a, b, v int) {
	if a <= l && r <= b {
		st.tree[p].setv = v
		st.tree[p].sumv = v * (r - l + 1)
		st.tree[p].addv = EMPTY
	} else {
		st.pushDown(p, r-l+1)
		m := (l + r) / 2
		if a <= m {
			st.Set(st.tree[p].l, l, m, a, b, v)
		}
		if m < b {
			st.Set(st.tree[p].r, m+1, r, a, b, v)
		}
		st.pushUp(p)
	}
}

func (st *SegmentTree) Query(p, l, r, a, b int) int {
	if a <= l && r <= b {
		return st.tree[p].sumv
	} else {
		st.pushDown(p, r-l+1)
		m := (l + r) / 2
		res := 0
		if a <= m {
			res += st.Query(st.tree[p].l, l, m, a, b)
		}
		if m < b {
			res += st.Query(st.tree[p].r, m+1, r, a, b)
		}
		st.pushUp(p)
		return res
	}
}
