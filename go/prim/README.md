## intro
prim和dijkstra很像，区别在于：
- prim把已经到达的点集视为一个点，然后进行松弛操作, d数组记录的是该点到点集的最小距离
- dijkstra的d数组记录的是起点到该点的距离

## 模板
- [cpp](./prim.cpp)
- [go](./prim.go)
