
## 同向双指针
大致思路为外循环固定移动j，内循环不断调整左区间使得判断条件成立 

```go
var i, j, cnt int 
for ; j<n; j++ {
    cnt++ 
    for ; i<j && (!check(i, j)); i++ {
        cnt-- 
    }
}
```

简单说一下，为什么双指针那么不好理解（不直观）。
代码实现上，看似我们主要移动右指针，被动的移动左指针，这使得我们疑惑，这样遍历是不是会遗漏某些可能？ 
但实际上，我们换个思路想，把左指针当成是主体，而右指针当成是从者。左指针一定把每个可能的位置都遍历了，而右指针恰好处在一个**不错**的位置（甚至仍然可以往右边移动）



参考题目：
- lc209 
- lc713 
- lc3 




## 相向双指针
[167. 两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/)


16. 最接近的三数之和 https://leetcode.cn/problems/3sum-closest/
18. 四数之和 https://leetcode.cn/problems/4sum/
611. 有效三角形的个数 https://leetcode.com/problems/valid-triangle-number/

