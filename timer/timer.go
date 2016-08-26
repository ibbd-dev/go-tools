package timer

import (
	"time"
)

type Entry struct {
	cmd      func()
	duration time.Duration
	time     time.Time
}

// 标识任务队列是否已经在运行中
var is_running = false

var tasks map[int]*Entry

func init() {
	tasks = make(map[int]*Entry)
}

// 增加任务函数
func AddFunc(cmd func(), duration time.Duration) {
	entry := &Entry{
		cmd:      cmd,
		duration: duration,
	}
	len_tasks := len(tasks)
	tasks[len_tasks] = entry
}

// 开始执行任务函数
// 注意: 精度不是很高, 可能会有1秒的延迟, 通常够用
func Start() {
	if is_running {
		return
	}
	is_running = true

	timer := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				now := time.Now()
				for key, entry := range tasks {
					if entry.time.IsZero() || now.After(entry.time.Add(entry.duration)) {
						tasks[key].time = now
						go entry.cmd()
					}
				}
			}
		}
	}()
}
