package timewheel

import (
	"fmt"
	"testing"
	"time"
)

var step, maxHour int 
var g = map[int]int64{
    50: 0,        // 50ms
    100: 0,       // 100ms 
    200: 0,       // 200ms
    1000: 0,      // 1s
    30000: 0,     // 30s 
    60000: 0,     // 1m 
    300000: 0,    // 5m
}

var cnt int 

func helper(interval int) {
    cnt += 1
    var d int64 
    now := time.Now().UnixMilli()
    if v := g[interval]; v != 0 {
        d = now - v - int64(interval)
        if d > int64(step) {                      // 判断误差
            fmt.Printf("now=%v, last=%v, delta=%v > step=%v\n", now, v, d, step)
            panic(interval)
        }
    }
    // fmt.Printf("pong every  %vms, millisecond=%v, delta=%v\n", interval, now, d) 
    g[interval] = now 
}


func TestTimeWheel(t *testing.T) {
    step = 50
    maxHour = 24
    tw, err := New(step, maxHour)
    if err != nil {
        panic(err)
    }
        
    for k, _ := range g {
        interval := k
        tw.Add(interval, func() {helper(interval)})
    }

    tw.Run()
}

func abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x 
}
