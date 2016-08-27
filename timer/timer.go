package timer

import (
	"time"
)

// 任务对象结构
type Entry struct {
	cmd      func()        // 任务函数
	time     time.Time     // 任务上次执行时间
	duration time.Duration // 任务执行的时间间隔
}

// 标识任务队列是否已经在运行中
var is_running = false

// 任务队列
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
		// 正在运行, 不需要再启动
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
