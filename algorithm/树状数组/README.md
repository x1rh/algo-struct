# intro
树状数组支持：
- 点更新
- 区间求和

# 含义
query(a) 维护的是[1, a]
query(b) 维护的是[1, b]
求[a, b] 等价求 query(b) - query(a-1)

# 注意事项
- 定义n为数组长度
- 数组下标从1到n，意味着数组长度为n+1
- query()的遍历的下界是1
- update的上界是n


# 模板
- [go](./BIT.go)
- [cpp](./BIT.cpp)

# 题目
- [LCR 170. 交易逆序对的总数](https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/) 
- [406. 根据身高重建队列](https://leetcode.cn/problems/queue-reconstruction-by-height/)
