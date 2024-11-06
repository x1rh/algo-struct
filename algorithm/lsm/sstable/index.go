package sstable

import (
	"encoding/json"
	"errors"
	"math"
	"os"

	"github.com/emirpasic/gods/trees/redblacktree"
)


type indexblock struct {
    redblacktree.Tree   // RedBlackTreeMap[string]indexElem 
}

type indexElem struct {
    Key string `json:"key"` 
    Offset int `json:"offset"` 
    Len    int `json:"len"` 
}


func newIndexBlock() *indexblock {
    idx := new(indexblock)
    idx.Tree = *redblacktree.NewWithStringComparator()
    return idx
}

// 找到 小于等于 key 的最右边的节点
// 简单包裹一层，方便替换底层数据结构，可以考虑写成接口
func (idx *indexblock) floor(key string) (*redblacktree.Node, bool){
    return idx.Floor(key)
}

func (idx *indexblock) load(fp *os.File, Offset, Len int64) error {
    if offset, err := fp.Seek(Offset, 0); err != nil {
        return err 
    } else if offset != Offset {
        return errors.New("offset != Offset")
    } 

    dump := make([]byte, Len)
    if n, err := fp.Read(dump); err != nil {
        return err 
    } else if int64(n) != Len {
        return errors.New("n != Len")
    }

    var data []indexElem 
    if err := json.Unmarshal(dump, &data); err != nil {
        return err 
    }
    
    idx.Clear()
    for _, el := range data {
       idx.Put(el.Key, el) 
    }

    return nil 
}

func (idx *indexblock) cal(n int) int {
    if n == 0 {
        return 0
    }
    return int(math.Ceil(math.Sqrt(float64(n))))            // todo:
}
