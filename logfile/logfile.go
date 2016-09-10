/*
程序运行日志，每分钟生成一个文件

格式如：
2016/08/09 14:08:55 [ERROR] hello world for error
@author Alex
@created_at 20160809
*/
package logfile

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

// 日志优先级
type Priority int

const (
	LEVEL_ALL Priority = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_OFF
)

var (
	//ip string // IP地址

	// 日志等级
	level_title = map[Priority]string{
		LEVEL_DEBUG: "DEBUG",
		LEVEL_INFO:  "INFO",
		LEVEL_WARN:  "WARN",
		LEVEL_ERROR: "ERROR",
		LEVEL_FATAL: "FATAL",
	}
)

type Logfile struct {
	Filename string   // 文件名
	Level    Priority // 日志记录等级

	f            *os.File    // 日志文件对象
	logger       *log.Logger // 日志对象
	is_json      bool        // 是否为json数据
	log_filename string      // 写入的文件名
	last_minute  string      // 最后一分钟
	write_mutex  sync.Mutex  // 写入锁
	switch_mutex sync.Mutex  // 切换文件锁
}

// json数据格式
type JsonData struct {
	Time time.Time   // 写入时间
	Type string      // 数据类型
	Data interface{} // 实际数据
}

// 切换文件
// 按分钟生成新的文件
func (lf *Logfile) switchFile() error {
	curr_minute := time.Now().Format("200601021504")
	if curr_minute == lf.last_minute {
		// 分钟数数没有改变则无需处理
		return nil
	}

	if lf.last_minute != "" {
		lf.f.Close()
	}

	// 组成新的文件名:给文件名加上分钟数的后缀
	lf.log_filename = lf.Filename + "." + curr_minute

	var err error
	lf.f, err = os.OpenFile(lf.log_filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	lf.last_minute = curr_minute

	if lf.logger != nil {
		lf.logger.SetOutput(lf.f)
	} else if lf.is_json {
		// json数据不需要输出前面的事件段
		lf.logger = log.New(lf.f, "", 0)
	} else {
		lf.logger = log.New(lf.f, "", log.LstdFlags)
	}

	return nil
}

// 写入一个字符串
func (lf *Logfile) Write(msg string) error {
	lf.switch_mutex.Lock()
	err := lf.switchFile()
	lf.switch_mutex.Unlock()
	if err != nil {
		return err
	}

	lf.write_mutex.Lock()
	lf.logger.Println(msg)
	lf.write_mutex.Unlock()

	return nil
}

// 设置信息写入等级
func (lf *Logfile) SetLevel(level Priority) {
	lf.Level = level
}

// 写入等级类数据
// 异步执行
func (lf *Logfile) WriteLevelMsg(msg string, log_level Priority) error {
	if log_level >= lf.Level && log_level > LEVEL_ALL && log_level < LEVEL_OFF {
		new_msg := "[" + level_title[log_level] + "]" + " " + msg
		go lf.Write(new_msg)
	}
	return nil
}

// 写入json数据
// 异步执行
func (lf *Logfile) WriteJson(data_type string, data interface{}) error {
	go func() {
		pack_data := &JsonData{
			Time: time.Now(),
			Type: data_type,
			Data: data,
		}

		bts, err := json.Marshal(pack_data)
		if err != nil {
			_log := GetAppLogfile(LEVEL_ERROR)
			_log.WriteLevelMsg("json.Marshal in writeJson", LEVEL_ERROR)
		}

		if lf.is_json == false {
			lf.is_json = true
		}
		err = lf.Write(string(bts))
		if err != nil {
			_log := GetAppLogfile(LEVEL_ERROR)
			_log.WriteLevelMsg(string(bts), LEVEL_ERROR)
		}
	}()

	return nil
}

func (lf *Logfile) Close() {
	lf.f.Close()
}
