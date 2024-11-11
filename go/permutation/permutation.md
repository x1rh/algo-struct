# Python
```python
from copy import deepcopy


def permutation(L, index, queue, answer):
    if index == len(L):
        answer.append(deepcopy(queue))
        return

    length = len(L[index])
    for i in range(0, length):
        queue.append(L[index][i])
        permutation(L, index+1, queue, answer)
        queue.pop()


if __name__ == '__main__':
    L = [['A'], ['B', 'C', 'D'], ['E'], ['F'], ['G', 'H']]
    answer = list()
    queue = list()
    permutation(L, 0, queue, answer)
    for each in answer:
        print(each)
```