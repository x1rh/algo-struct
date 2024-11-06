# intro
- 线段树（基础）
- 动态开点线段树
- 权值线段树
- 可持久化线段树


# 线段树使用场景
- 点更新
- 点查询
- 区间加x
- 区间置x
- 区间查询
- 区间中最大连续子数组和 [leetcode-53](https://leetcode.cn/problems/maximum-subarray/)
    - iSum = l.iSum + r.iSum
    - lSum = max(l.lSum, l.iSum + r.lSum)
    - rSum = max(r.rSum, r.iSum + l.rSum)
    - mSum = max(max(l.mSum, r.mSum), l.rSum + r.lSum)
- 判断新插入区间和已有区间是否重叠
- 快速统计前缀和（因为有更新，所以前缀和需要动态更新）
- 线段树最底层全是1或0解决某些问题
- 线段的交/区间的交：判断是否和已存在的区间相交
- 权值线段树：
    - 1）求数值k是第几大 
    - 2）二分求第k大
- 求满足`pre[r]-pre[l] < t` 的二元组`(l, r)`的个数
- 求逆序数（二维偏序问题）
    - 普通线段树的做法：
        - 对待排序序列做离散化，映射到区间`[1, n]`上
        - 区间`[1,n]`每个值初始化为0，表示没有插入
        - 按顺序插入序列，对新插入项，其在区间`[1,n]`中排k，则求区间`[k+1, n]`中插入值的个数cnt_k。
        - cnt_k 即当前插入项目对逆序对个数的贡献
    - 权值线段树的做法
- 原数组a[i]中，使得 `a[i]>=k` 的最小的`i`
    - 使用线段树，是因为a[i]被动态更新
    - 维护区间最大值


# 使用线段树时的其他问题
- 离散化
    - 数值过大时需要离散化
    - 离散化需要离线
    - 在线问题无法离散化？答：动态开点线段树
- 建树
- 标记
    - 有时候标记可以和标记值共用一个变量
    - 如果不行，则需要新开一个布尔变量


# 线段树的常数
自顶向下线段树常数较大，另有常数较小的[zkw线段树](https://zhuanlan.zhihu.com/p/361935620)


# 线段树编写注意问题
- pushdown最容易出问题
- 初始化容易出问题




# 线段树模板-golang
- 单点更新、区间查询模板 [代码](./单点更新&区间查询.go)
- 区间加x、区间查询
- 区间置x、区间查询
- 混合：区间加x、区间置x，区间查询
    - 同时存在setv和addv两种标记怎么处理？ 
        - 这种情形属于先set后add导致的
        - 所以在pushdown的时候，先处理setv标记 


# 动态开点线段树
- 动态线段树写成node和tree两个结构体这种形式，缓存命中率更高（因为数据在一块）
- 想象树的节点是虚拟的，只有在使用时再把它变为真实的。
- 什么时候化虚为实？pushdown的时候，这时候可以检查处儿子节点是否存在，不存在则添加一个。
- 注意点
    - 节点数量计算
    - 注意根节点初始化
    - 注意pushdown中新建节点的初始化
    - 注意空标记
    - 注意左右儿子的表示方法发生了变化
    - 推荐使用指针的方法进行实现，实在想偷懒用vector，不要开静态数组。

## 节点计算
节点个数=查询次数 * log2(值域)    

举例说明：
- [codeforeces1042D](https://codeforces.com/contest/1042/problem/D) 这题查询次数20w， 值域[-2e14, 2e14]，节点个数=`20e4 * log2(4e14) = 20e4 * 49 = 9800000`
- [leetcode729](https://leetcode-cn.com/problems/my-calendar-i/) 这题查询次数1000，值域[0, 1e9], 问题来了，为什么这题的节点个数是 2 * 1000 * log2(1e9) ?

## 动态开节点注意事项
按理说一个题目如果能用线段树做，那么节点个数一定能开出来，但是在实现上，如果处理得不好，那么就容易RE（RE还报WA）

实现上，如果update()中，无脑在pushdown时开左右两个儿子节点，那么可能会造成非常大的开销。为了改进这一点，必须做到只有用到某一个儿子时，再动态给它开一个点。

负区间问题：
(-1 + 0) / 2 = 0 
求mid时，一定要写成l + (r-l)/2 


## 看看以下两题的pushdown中开点的做法
如果例题1使用例题2的建点方法，内存不够，原因存在大量的空节点
所以最好还是离散化+动态开点?

- 例题1：[codeforeces1042D](https://codeforces.com/contest/1042/problem/D) - [代码](../../codeforces/1042D-sol1.cpp) 
- 例题2：[leetcode729](https://leetcode-cn.com/problems/my-calendar-i/) - [代码](../../leetcode/0729(我的日程安排表I).md) 
- 例题3：[6066-统计区间中的整数数目](../leetcode/6066(统计区间中的整数数目).md)




# 权值线段树
- 权值线段树类似于桶`cnt[]`，`cnt[x]`表示值为x的数有`cnt[x]`个
- 一般使用动态开点的方式实现
- 可以解决两类问题:
    - 求数值k在桶中排第几名，对应实现是`rank(k)`
    - 二分求第k名的值是什么，对应实现是`kth(k)`
- 权值线段树用动态开点线段树可以很自然的实现


下面的代码片段来自：https://zhuanlan.zhihu.com/p/492995124
```cpp
struct node {
    int sum;
    int l, r;
}tree[MAXN << 2];
#define ls (2*rt)
#define rs (2*rt + 1)
int a[MAXN];
int n, m;
void push_up(int rt) {
    tree[rt].sum = tree[ls].sum + tree[rs].sum;
}
void build(int rt, int L, int R) {
    tree[rt].l = L, tree[rt].r = R;
    if (L == R) {
        tree[rt].sum = a[L];
        return;
    }
    int mid = (L + R) / 2;
    build(ls, L, mid);
    build(rs, mid + 1, R);
    push_up(rt);
}
void add(int rt, int pos, int x) {
    if (tree[rt].l == tree[rt].r) {
        tree[rt].sum += x;
        return;
    }
    int mid = (tree[rt].l + tree[rt].r) / 2;
    if (pos <= mid)add(ls, pos, x);
    else add(rs, pos, x);
    push_up(rt);
}
int que(int rt, int L, int R) {
    if (tree[rt].r < L || tree[rt].l > R)return 0;
    if (L <= tree[rt].l && tree[rt].r <= R)return tree[rt].sum;
    return que(ls, L, R) + que(rs, L, R);
}
int kth(int rt,int k) {
    if (tree[rt].l == tree[rt].r)return tree[rt].l;
    if (tree[ls].sum >= k)return kth(ls, k);
    else return kth(rs, k - tree[ls].sum);
}
int _rank(int k) {
    return que(1, 1, k - 1) + 1;
}
void insert(int x) {
    add(1, x, 1);
}
void del(int x) {
    add(1, x, -1);
}
void slove() {
    build(1, 1, 100000);
    vector<int>v = { 1,3,4,4,4,7,9,9,12 };
    for (int x : v)insert(x);
    cout << kth(1, 4) << endl;
    cout << kth(1, 6) << endl;
    cout << kth(1, 7) << endl;
    cout << _rank(4) << endl;
    cout << _rank(9) << endl;
    cout << _rank(12) << endl;   
}
```



# 可持久化线段树
- 对树上存在修改的点，全部开新点处理，复用未修改的点。
- 可持久化线段树准确上来说不是树




