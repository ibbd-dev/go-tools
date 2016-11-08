package timer

import (
	//"fmt"
	"time"
)

// 任务对象结构
type Entry struct {
	cmd           func()        // 任务函数
	duration      time.Duration // 任务执行的时间间隔
	lastBeginTime time.Time     // 任务上次开始时间
	lastEndTime   time.Time     // 任务上次结束时间
}

// 任务队列
var tasks []*Entry

func init() {
	tasks = make([]*Entry, 0, 8)
	start()
}

// 增加任务函数
func AddFunc(cmd func(), duration time.Duration) {
	entry := &Entry{
		cmd:      cmd,
		duration: duration - time.Millisecond, // 减少每次运行时间的误差
	}
	tasks = append(tasks, entry)
}

// 获取现在时间的函数
var nowFunc = time.Now

const (
	// 时钟的周期
	Duration = time.Millisecond * 200
)

// 开始执行任务函数
func start() {
	timer := time.NewTicker(Duration)
	go func() {
		for {
			select {
			case <-timer.C:
				for _, entry := range tasks {
					go runTask(entry)
				}
			}
		}
	}()
}

func runTask(entry *Entry) {
	//println(entry.duration)
	//println(entry.lastBeginTime.Format(time.RFC3339), entry.lastEndTime.Format(time.RFC3339))
	if !entry.lastBeginTime.IsZero() && entry.lastEndTime.IsZero() {
		// is running
		//println("return1")
		return
	}

	if entry.lastBeginTime.After(entry.lastEndTime) {
		// is running
		//println("return2")
		return
	}

	if entry.lastBeginTime.IsZero() || nowFunc().After(entry.lastBeginTime.Add(entry.duration)) {
		entry.lastBeginTime = nowFunc()
		entry.cmd()
		entry.lastEndTime = nowFunc()
		return
	}
}
