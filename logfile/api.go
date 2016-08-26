package logfile

import (
	"git.ibbd.net/dsp/go-config"
)

var (
	app_runtime_log *Logfile // 全局的程序运行日志对象
	bid_json_log    *Logfile // 竞价接口json数据对象
	event_json_log  *Logfile // 事件接口json数据对象
)

func init() {
	app_runtime_log = &Logfile{
		Filename: goConfig.PathAppRuntimeLog,
		Level:    LEVEL_ERROR,
	}

	bid_json_log = &Logfile{
		Filename: goConfig.PathBidDataLog,
	}

	event_json_log = &Logfile{
		Filename: goConfig.PathEventDataLog,
	}
}

// 写入程序运行日志
func GetAppLogfile(level Priority) *Logfile {
	app_runtime_log.SetLevel(level)
	return app_runtime_log
}

// 写入竞价接口的json数据
func GetBidLogfile() *Logfile {
	return bid_json_log
}

// 写入事件接口的json数据
func GetEventLogfile() *Logfile {
	return event_json_log
}
