package command

import "unsafe"

const (
	OpPut = 1
	OpDel = 2
	OpGet = 4
)

type Command struct {
	Op  int    `json:"op"` // 1: set, 2: del
	Key string `json:"key"`
	Val string `json:"val"`
}

// 一个不太考究的内存计算公式
func (c *Command) Size() int {
    size := int(unsafe.Sizeof(*c))
    size += len(c.Key) * 2 + len(c.Val)   // 考虑到TreeMap[key] = Command{op, key, val}, 把key也计算在内
	return size 
}
