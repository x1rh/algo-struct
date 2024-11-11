package timewheel

import (
	"container/list"
	"errors"
	"log"
	"sync"
	"time"
)

const (
    Hour_MS   = 1000 * 60 * 60 
    Minute_MS = 1000 * 60
    Second_MS = 1000 
)

var ErrInvalidStep = errors.New("invalid step")
var ErrInvalidInterval = errors.New("invalid interval")
var ErrFatal = errors.New("fatal logic error")

// timeinfo 用来记录当前时间轮在哪一个slot 
// 例如22:23:33:800, step 为20ms，那么表示为timeinfo{ms:40, s: 33, m: 23, h: 22}
type timeinfo struct {
    ms int   // 1000ms / step 
    s  int   
    m  int 
    h  int 
}

type Event struct {
    id int            // 
    interval int      // 事件每隔interval毫秒运行一次 
    time timeinfo     // 记录上一次运行的时间信息, 配合step 和 interval 可以计算出下一次运行的具体时间 
    callback func()   // 事件的运行函数
}

type TimeWheel struct {
    millisecondSlotCnt int   
    secondSlotCnt      int    
    minuteSlotCnt      int    
    hourSlotCnt        int   

    step     int         // step控制精度，即把1000毫秒划分成x份，每份长度为step
    eventCnt int         // 时间轮中总共有多少event, 用作eventId

    time   timeinfo 
    slots  []*list.List  // from 0 ~ n - 1 : millisecond | second | minute | hour 
    locker *sync.Mutex
}

// maxHour 指定最多多少个小时 
func New(step, maxHour int) (*TimeWheel, error) {
    if 1000 % step != 0 {
        return nil, ErrInvalidStep
    } 
    tw := &TimeWheel {
        step: step, 
        millisecondSlotCnt: 1000 / step,  // 1秒划分成多少个slot
        secondSlotCnt: 60,                // 60个1秒的slot 
        minuteSlotCnt: 60,                // 60个1分的slot
        hourSlotCnt: maxHour,             // maxHour个1小时的slot
        locker: &sync.Mutex{},
    }
    tw.slots = make([]*list.List, tw.millisecondSlotCnt + tw.secondSlotCnt+ tw.minuteSlotCnt + tw.hourSlotCnt)
    for i := range tw.slots {
        tw.slots[i] = list.New()
    }
    return tw, nil 
}

// Add 添加Event到时间轮中
// 每次插入的 Event，都插入到当前时间轮时间的 interval 毫秒后
func (tw *TimeWheel) Add(interval int, callback func()) error {
    if interval < tw.step || interval % tw.step != 0 || 
        interval >= tw.step * tw.millisecondSlotCnt * tw.secondSlotCnt * tw.minuteSlotCnt * tw.hourSlotCnt {
        return ErrInvalidInterval
    }
    e := &Event {
        id: tw.nextEventId(),
        interval: interval,
        time: tw.time,
        callback: callback,
    }
    tw.locker.Lock()
    defer tw.locker.Unlock()
    tw.insertAfter(e.interval, e)
    return nil 
}

// Run 主循环
func (tw *TimeWheel) Run() {
    for { 
        time.Sleep(time.Millisecond * time.Duration(tw.step))      
        last := tw.time 
        tw.time = tw.afterInterval(tw.step)

        var index int            
        tw.locker.Lock() 
        if tw.time.h != last.h {
            index = tw.time.h + tw.minuteSlotCnt + tw.secondSlotCnt + tw.millisecondSlotCnt 
        } else if tw.time.m != last.m {
            index = tw.time.m + tw.secondSlotCnt + tw.millisecondSlotCnt
        } else if tw.time.s != last.s {
            index = tw.time.s + tw.millisecondSlotCnt 
        } else if tw.time.ms != last.ms {
            index = tw.time.ms 
        } else {
            panic(ErrFatal)
        }

        if err := tw.handle(index); err != nil {
            log.Println("handle err: ", err)
        }
        tw.locker.Unlock()
    }
}

// handle 处理index指定的slot, 每隔slot里存储着一个事件链表
// 处理到某个slot，意味着这个slot里的event，要么现在执行，要么往下推
func (tw *TimeWheel) handle(index int) error {
    for it := tw.slots[index].Front(); it != nil; it = it.Next() {
        event := it.Value.(*Event)
        now := getMs(tw.step, tw.time)
        last := getMs(tw.step, event.time)
        if event.interval == now - last {
            event.callback()
            event.time = tw.time 
            tw.insertAfter(event.interval, event)
        } else {
            tw.insertAfter(last + event.interval - now, event)        // last + event.interval 一定大于 now
        }
    }
    tw.slots[index].Init()
    return nil 
}

// insert 将事件e插入到当前时间轮时间的interval毫秒后
// 注意不要修改e的timeinfo
func (tw *TimeWheel) insertAfter(interval int, e *Event) {
    future := tw.afterInterval(interval)
    var index int 
    if future.h != tw.time.h {
        index = tw.millisecondSlotCnt + tw.secondSlotCnt + tw.minuteSlotCnt + future.h
    } else if future.m != tw.time.m {
        index =  tw.millisecondSlotCnt + tw.secondSlotCnt + future.h  
    } else if future.s != tw.time.s {
        index = tw.millisecondSlotCnt + future.s  
    } else if future.ms != tw.time.ms {
        index = future.ms
    }
    tw.slots[index].PushBack(e)
}


// 计算当前时间轮经过interval毫秒后的时间 
func (tw *TimeWheel) afterInterval(interval int) timeinfo {
    future := getMs(tw.step, tw.time) + interval 
    return timeinfo{
        h:  future / Hour_MS,
        m:  (future % Hour_MS) / Minute_MS, 
        s:  (future % Minute_MS) / Second_MS,
        ms: (future % Second_MS) / tw.step, 
    }
}

func (tw *TimeWheel) nextEventId() (id int){
    id = tw.eventCnt
    tw.eventCnt += 1
    return 
}


// 计算指定step和timeinfo所代表的时间，单位是毫秒
func getMs(step int, t timeinfo) int {
    return step * t.ms + t.s * Second_MS + t.m * Minute_MS + t.h * Hour_MS 
}
