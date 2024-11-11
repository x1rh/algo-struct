package dijkstra

const INF = 0x3f3f3f3f

type node struct {
	y, c int
}

// dijkstra 计算单源最短路，起点是start，节点标号从0至n-1
func Dijkstra(g map[int][]node, start, n int) []int {
	d := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		d[i] = INF
	}

	d[start] = 0
	for i := 0; i < n-1; i++ {
		min := INF
		idx := -1
		for j := 0; j < n; j++ {
			if !vis[j] && d[j] < min {
				min = d[j]
				idx = j
			}
		}

		if idx == -1 {
			return d // 意味着没有节点可以进行松弛操作，dijkstra结束
		}

		vis[idx] = true
		v := g[idx]
		m := len(v)
		for j := 0; j < m; j++ {
			to := v[j].y
			cost := v[j].c
			if d[to] > d[idx]+cost {
				d[to] = d[idx] + cost
			}
		}
	}

	return d
}

// leetcode-743
func networkDelayTime(times [][]int, n int, k int) int {
	g := make(map[int][]node)
	for i := 0; i < len(times); i++ {
		g[times[i][0]-1] = append(g[times[i][0]-1], node{times[i][1] - 1, times[i][2]})
	}
	d := dijkstra(g, k-1, n)
	ans := -1
	for i := 0; i < n; i++ {
		if d[i] == INF {
			ans = -1
			break
		} else {
			if d[i] > ans {
				ans = d[i]
			}
		}
	}

	return ans
}
