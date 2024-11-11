type node struct {
	sumv int
    maxv int 
	l, r int 
	addv, setv int 
	addf, setf bool
}

type segmenttree struct{
    tree []node 
}

func (seg *segmenttree) pushUp(p int) {
	o := &seg.tree[p] 
	o.sumv = seg.tree[o.l].sumv + seg.tree[o.r].sumv 	
    o.maxv = max(seg.tree[o.l].maxv, seg.tree[o.r].maxv)
}

func (seg *segmenttree) pushDown(p, Len int) {
	o := &seg.tree[p] 
	if o.l == 0 {
		o.l = len(seg.tree)
		seg.tree = append(seg.tree, node{})
        o = &seg.tree[p]        
	}
	if o.r == 0 {
		o.r = len(seg.tree)
		seg.tree = append(seg.tree, node{})
        o = &seg.tree[p] 
	}

	rl := Len / 2  
	ll := Len - rl

	if o.setf {
		seg.doSet(o.l, ll, o.setv)
		seg.doSet(o.r, rl, o.setv)
		o.setf = false 
	}
	if o.addf {
		seg.doAdd(o.l, ll, o.addv)
		seg.doAdd(o.r, rl, o.addv)
		o.addf = false 
	}
}

func (seg *segmenttree) doSet(p, Len, setv int) {
	o := &seg.tree[p]
	o.setv = setv 
	o.setf = true 
    o.addv = 0 
	o.addf = false 
	o.sumv = setv * Len 
    o.maxv = setv 
}

func (seg *segmenttree) doAdd(p, Len, addv int) {
	o := &seg.tree[p]
	o.addv += addv 
    o.maxv += addv 
	o.sumv += addv * Len 
	o.addf = true 
}

func (seg *segmenttree) Add(p, l, r, ql, qr, v int) {
	if ql <= l && r <= qr {
		seg.doAdd(p, r-l+1, v)
	} else {
		m := (l + r) / 2 
		seg.pushDown(p, r-l+1)
		if ql <= m {
			seg.Add(seg.tree[p].l, l, m, ql, qr, v)
		}
		if qr > m {
			seg.Add(seg.tree[p].r, m+1, r, ql, qr, v)
		}
		seg.pushUp(p)
	}
}

func (seg *segmenttree) Set(p, l, r, ql, qr, v int) {
	if ql <= l && r <= qr {
		seg.doSet(p, r-l+1, v)
	} else {
		m := (l + r) / 2
		seg.pushDown(p, r-l+1)
		if ql <= m {
			seg.Set(seg.tree[p].l, l, m, ql, qr, v)
		}
		if qr > m {
			seg.Set(seg.tree[p].r, m+1, r, ql, qr, v)
		}
		seg.pushUp(p)
	}
}

func (seg *segmenttree) Query(p, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return seg.tree[p].sumv 
	} else {
		m, res := (l + r) / 2, 0
		seg.pushDown(p, r-l+1)
		if ql <= m {
			res += seg.Query(seg.tree[p].l, l, m, ql, qr)
		} 
		if qr > m {
			res += seg.Query(seg.tree[p].r, m+1, r, ql, qr)
		}
		seg.pushUp(p)
		return res 
	}
}

func (seg *segmenttree) QueryMax(p, l, r, ql, qr int) int {
    if ql <= l && r <= qr {
        return seg.tree[p].maxv 
    } else {
        m, res := (l+r)/2, 0 
        seg.pushDown(p, r-l+1)
        if ql <= m {
            res = max(res, seg.QueryMax(seg.tree[p].l, l, m, ql, qr))
        }
        if qr > m {
            res = max(res, seg.QueryMax(seg.tree[p].r, m+1, r, ql, qr))
        }
        seg.pushUp(p)
        return res 
    }
}

func NewSegmentTree() *segmenttree {
	return  &segmenttree{
        tree: make([]node, 1),
    }
}

func max(a, b int) int {
    if a > b {
        return a 
    }
    return b 
}