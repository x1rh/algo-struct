## 快速访问

- 同余运算
- 费马小定理
- 快速幂
- 矩阵快速幂
- 素性判定
  - 试除法
  - 费马小定理素性判定
  - Miller-Rabin 素性判定
- 素数筛法
  - 暴力
  - 埃式筛
  - 欧拉筛
- 质因数分解
  - 暴力
  - Pollard-Rho

## 同余运算

加法原理:  
(a+b)%n = [(a%n) + (b%n)] % n

乘法原理:  
(a*b)%n = [(a%n) * (b%n)] % n

## 费马小定理

- [同余运算和费马小定理的证明 - 黄兢成的文章 - 知乎](https://zhuanlan.zhihu.com/p/75685377)

## 快速幂

- a^n
- a^n % p

## 矩阵快速幂

## 1. 素数判定

### 1.1 试除法

参考`is_prime.go`

时间复杂度: O(sqrt(n))

### 1.2 Fermat 素性判定

详细见`fermat.go`

时间复杂度: O(k log n), k 是测试次数

费马小定理的逆命题不成立

测试次数如何确定？

### 1.3 Miller-Rabin 素性判定

## 2. 筛法

### 2.1 暴力

### 2.2 埃式筛

时间复杂度：O(log(log n))
缺点：同一个数可能被筛多次，例如 12 被 2 和 3 筛两次

### 2.3 欧拉筛

时间复杂度：O(n)

## 3. 质因数分解

### 3.1 Pollard-Rho

## 测试

单核测试，一些规模大的测试运行很慢(todo: 可以均分然后并行运行，以利用多核)

```shell
go test -v
```

## 参考资料

- [Pecco 算法学习笔记(17): 素数筛](https://zhuanlan.zhihu.com/p/100051075)
- [算法学习笔记(48): 米勒-拉宾素性检验](https://zhuanlan.zhihu.com/p/220203643)
- [Pecco 算法学习笔记(55): Pollard-Rho 算法](https://zhuanlan.zhihu.com/p/267884783)
- [肖有量 分解质因数-Pollard‘s Rho](https://blog.csdn.net/qq_43449564/article/details/123979433)
