# Intro

## ABA 问题

线程 1 读到 A， 此时线程 2 将 A 改为了 B 后，它（或者其他的线程）又将其改回了 A 。
此时线程 1 又读到 A，并认为其没有发生变化。

# ListQueue

参考 [Simple, Fast, and Practical Non-Blocking and Blocking Concurrent Queue Algorithms - Maged M. Michael Michael L. Scot](https://www.cs.rochester.edu/u/scott/papers/1996_PODC_queues.pdf)

通过引用计数的方式解决 ABA 问题。

# ArrayQueue

# 其他实现
