type node struct {
	l, r *node 
	sumv int
	addv int 
	addf bool 
}

type segmenttree struct {
	root *node 
}

func pushUp(o *node) {
	var l, r int 
	if o.l != nil {
		l = o.l.sumv 
	}	
	if o.r != nil {
		r = o.r.sumv 
	}
	o.sumv = l + r 
}

func doAdd(o *node, Len, addv int) {
	o.addv += addv 
	o.sumv += Len * addv 
	o.addf = true 
}

func pushDown(o *node, Len int) {
	if o.addf {
		if o.l == nil {
			o.l = &node{}
		}
		if o.r == nil {
			o.r = &node{} 
		}
		rl := Len / 2
		ll := Len - rl 
		doAdd(o.l, ll, o.addv)
		doAdd(o.r, rl, o.addv)
		o.addf = false 
		o.addv = 0
	}
}

func (seg *segmenttree) Add(o *node, l, r, ql, qr, v int) {
    if r < ql || l > qr {                  // 这个其实是防止错误调用，如果[ql, qr]是[l,r]的子集，那么这个判断不是必须的
        return 
    }
	if ql <= l && r <= qr {
		doAdd(o, r-l+1, v)
	} else {
		pushDown(o, r-l+1)
		m := l + (r-l) / 2
		if ql <= m {
            if o.l == nil {
                o.l = &node{} 
            }
			seg.Add(o.l, l, m, ql, qr, v)
		}
		if qr > m {
            if o.r == nil {
                o.r = &node{}
            }
			seg.Add(o.r, m+1, r, ql, qr, v)
		}
		pushUp(o)
	}
}

func (seg *segmenttree) Query(o *node, l, r, ql, qr int) int {
    if o == nil {
        return 0
    }
	if ql <= l && r <= qr {
		return o.sumv 
	} else {
		pushDown(o, r-l+1)
		m, res := l + (r-l) / 2, 0
		if ql <= m {
			res += seg.Query(o.l, l, m, ql, qr)
		}
		if qr > m {
			res += seg.Query(o.r, m+1, r, ql, qr)
		}
		pushUp(o)
		return res 
	}
}

func NewSegmentTree() *segmenttree {
	return &segmenttree{root: &node{}}
}