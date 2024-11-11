package disjoint_set

// V1: 路径压缩
type UnionFindSetV1 struct {
	fa []int
}

func Init(s *UnionFindSetV1, n int) {
	s.fa = make([]int, n)
	for i := 0; i < n; i++ {
		s.fa[i] = i
	}
}

func (s *UnionFindSetV1) find(x int) int {
	if x == s.fa[x] {
		return x
	} else {
		s.fa[x] = s.find(s.fa[x])
		return s.fa[x]
	}
}

func (s *UnionFindSetV1) union(x, y int) bool {
	px := s.find(x)
	py := s.find(y)
	if px != py {
		s.fa[px] = py
		return true
	} else {
		return false
	}
}
