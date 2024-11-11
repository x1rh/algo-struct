package lsm

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)


// 正确性验证：
// 随机产生1000W个数作为key，随机选择1个数随机的在PUT、GET、DEL这三个操作中选择一个，重复1亿次
// PUT的时候随机生成对应的val
func TestLSMTree(t *testing.T) {
    lsm, err := NewLSMTree(Threshold(4 * 1024 * 1024))     // 4 MB
    if err != nil {
        t.Fatal(err)
    }
    n := 1000000
    var a []string 
    g := make(map[string]string)
    rand.Seed(time.Now().Unix())
    var putOpCnt, getOpCnt int 
    for i:=0; i<n; i++ {
        op := rand.Intn(3)
        k := strconv.Itoa(rand.Int()) 
        if op == 0 {                           
            v := strconv.Itoa(rand.Int())
            g[k] = v 
            a = append(a, k)
            lsm.Put(k, v) 
            putOpCnt += 1
        } else {
            if op == 2 {
                if len(a) == 0 {
                    continue 
                }
                k = a[rand.Intn(len(a))]
            }
            v1, ok1 := lsm.Get(k)
            v2, ok2 := g[k]
            if ok1 == ok2 && v1 == v2 {
                continue 
            } else {
                t.Fatalf("k=%s, v1=%s, ok1=%v, v2=%s, ok2=%v", k, v1, ok1, v2, ok2)
            }
            getOpCnt += 1
        }
        if i % 1000 == 0 {
            t.Log(i)
        }
    }
    t.Logf("OpPutCnt=%d, OpGetCnt=%d\n", putOpCnt, getOpCnt)
}



func TestClearAll(t *testing.T) {

}
