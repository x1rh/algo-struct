# intro

LSMTree (log structured merge tree)

SSTable 的存储：

1. 首先 SSTable 是一个有序的列表，存储在磁盘上时也是如此
2. 其次为了加速在文件中的查找，在存储到磁盘时，加入一个查找索引表，用于快速查找 key 大致位置
3. SSTable 在内存中只存储一个较小的查找索引表, 而不是整个键值对表
4. 实践中有给每个 SSTable 设置一个布隆过滤器的，用来快速判断 key 是否存在于当前 SSTable

两个恢复场景：

1. 当前 LSMTree 发生某个异常，可以通过当前的 wal 恢复 memtable
2. 之前未持久化的 SSTable 发生某个异常，通过之前的 wal 继续完成 SSTable 的持久化

坑：
假设疯狂对同一个 key 进行 put 操作，那么将会在 wal 中产生大量的记录，而 memtable 的大小不发生变化
所以持久化的标准，需要同时考虑.wal 的大小和 memtable 的大小

out-place update, 写快读慢

特点：写快、读慢
适用场景：多写少读

# 缺点/问题

- 空间放大
- 读放大
- 写放大
