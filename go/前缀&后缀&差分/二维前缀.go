type NumMatrix struct {
    g [][]int 
}

func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    g := make([][]int, m+1)
    for i:=0; i<=m; i++ {
        g[i] = make([]int, n+1)
    }
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            g[i+1][j+1] = g[i+1][j] + g[i][j+1] - g[i][j] + matrix[i][j] 
        }
    }
    return NumMatrix{g: g} 
}


// leetcode 304. 二维区域和检索 - 矩阵不可变
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.g[row2+1][col2+1] - this.g[row1][col2+1] - this.g[row2+1][col1] + this.g[row1][col1]
}
