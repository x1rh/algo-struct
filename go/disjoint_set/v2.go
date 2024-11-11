package disjoint_set

// V2: 路径压缩 & 按秩合并
type UnionFindSetV2 struct {
	fa, rank []int
}

func Init(s *UnionFindSetV2, n int) {
	s.fa = make([]int, n)
	s.rank = make([]int, n)
	for i := 0; i < n; i++ {
		s.fa[i] = i
		s.rank[i] = 1
	}
}

// 路径压缩，将返回值置为查询值x的父亲
func (s *UnionFindSetV2) find(x int) int {
	if s.fa[x] == x {
		return x
	} else {
		s.fa[x] = s.find(s.fa[x])
		return s.fa[x]
	}
}

// 按秩压缩，将秩小的合并到大的上面
// 已经联通，返回false
// 可以联通，返回true
func (s *UnionFindSetV2) union(x, y int) bool {
	px := s.find(x)
	py := s.find(y)
	if px != py {
		if s.rank[px] <= s.rank[py] {
			s.fa[px] = py
			if s.rank[px] == s.rank[py] {
				s.rank[py] += 1
			}
		} else {
			s.fa[py] = px
		}
		return true
	} else {
		return false
	}
}
