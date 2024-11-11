package kruskal


type UnionFindSet struct {
    n int 
    fa, rk []int 
}

func New(n int) *UnionFindSet {
    s := &UnionFindSet{n: n, fa: make([]int, n), rk: make([]int, n)}
    for i:=0; i<n; i++ {
        s.fa[i] = i 
        s.rk[i] = 1
    }
    return s 
}

func (s *UnionFindSet) Find(x int) int {
    if s.fa[x] == x {
        return x 
    } else {
        s.fa[x] = s.Find(s.fa[x])
        return s.fa[x]
    }
}

func (s *UnionFindSet) Union(x, y int) bool {
    fx := s.Find(x)
    fy := s.Find(y)
    if fx == fy {
        return false 
    } else {
        if s.rk[fx] <= s.rk[fy] {
            s.fa[fx] = fy 
            if s.rk[fx] == s.rk[fy] {
                s.rk[fx] += 1
            }
        } else {
            s.fa[fy] = fx
        }
        return true 
    }
}

// 检查是否联通
func (s *UnionFindSet) check() bool {
    root := s.Find(0)
    for i:=0; i<s.n; i++  {
        if s.Find(i) != root {
            return false 
        }
    }
    return true 
}

func kruskal(c [][]int, n int) int {
    var res int   // 最小生成树的值
    s := New(n)
    for _, e := range c {
        if s.Union(e[0]-1, e[1]-1) {
            res += e[2]
        }
    }
    if s.check() {
        return res 
    } else {
        return -1 
    }
}

