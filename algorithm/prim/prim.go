package prim 

type pair struct {
    to, d int 
}

type G struct {
    g [][]pair 
    n int 
}

func (g *G) prim(start int) int{
    inf := 0x3f3f3f3f 
    dist := make([]int, g.n)
    for i:=0; i<g.n; i++ {
        dist[i] = inf 
    }
    vis := make([]bool, g.n)
    dist[start] = 0 
    res := 0

    for i:=0; i<g.n; i++ {
        to := -1 
        min := inf 
        for j:=0; j<g.n; j++ {
            if !vis[j] && dist[j] < min {
                to = j
                min = dist[j] 
            }        
        }

        if to == -1 {
            // 不连通
        }

        res += min 
        vis[to] = true 
        for j:=0; j<len(g.g[to]); j++ {
            u := g.g[to][j].to
            d := g.g[to][j].d 
            if !vis[u] && dist[u] > d {
                dist[u] = d 
            } 
        }
    }
    return res
}
