# 前缀
- 定义原数组`a[i]`, 0-indexed
- 定义前缀和数组`pre[i]`, 0-indexed
    - 范围是0到n，共n+1个数。 
    - `pre[i] = sum(a[0], ..., a[i-1])`
    - `pre[0]=0`
    - `pre[r]-pre[l] = sum(a[l], ..., a[r-1])`

# 二维前缀 
[leetcode 304. 二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/description/) 

# 后缀 

# 一维差分

参考： [那些小而美的算法技巧：前缀和/差分数组 - labuladong的文章 - 知乎](https://zhuanlan.zhihu.com/p/301509170) 

需要理解+incr和-incr的含义 

此外，在实际编写代码时，通常多开1，省去边界判定 


leetcode370
leetcode1109
leetcode1094


# 二维差分 

leetcode 6292. 子矩阵元素加 1