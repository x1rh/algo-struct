# intro
- set维护一个两两不想交的区间集合，每个元素为元组`(r, l)`
- 维护(r,l)而不是(l,r)的原因是，方便使用`pair`和`lower_bound()`
- 新加入一个区间，那么把所有涉及的区间进行合并，删除被合并的区间，新加入合并后的区间。

# 代码
[leetcode 6066](https://leetcode.cn/problems/count-integers-in-intervals/)
```cpp
class CountIntervals {
public:
    typedef pair<int, int> pii;    
    set<pii> s;
    int cnt;

    CountIntervals() {
        cnt = 0;
    }
    
    void add(int left, int right) {
        int l = left;
        int r = right;
        auto it = s.lower_bound(pii(left-1, -1));
        while (it != s.end()) {
            if (it->second > right + 1) break;
            l = min(l, it->second);
            r = max(r, it->first);
            cnt -= it->first - it->second + 1;
            s.erase(it++);
        }
        s.insert(pii(r, l));
        cnt += r - l + 1;
    }
    
    int count() {
        return cnt;
    }
};
```
1. 找到第一个区间`(i_l,i_r)`，使得 `i_r>=left-1`。 举例说明：`[1,2], [3, 4]`
2. 设元组(l, r)为最终插入值，那么不断更新这两个值即可。
3. 退出条件：类似于1， 找到第一个`(j_l, j_r)`, 使得`j_l-1>right`， 即`right+1<j_l`时退出，例如`[1,2], [4, 5]`

## 题目
- [lc56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
    - 简单的区间合并（无顺序要求）
- [lc6066](https://leetcode.cn/problems/count-integers-in-intervals/)
- [lc715 range模块]()
    - 提供了区间的插入合并、区间删除、区间查询：[leetcode-715题解](../leetcode/0715(Range模块).md)


