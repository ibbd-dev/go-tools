package logfile

import (
	"fmt"
	"testing"
	"time"
)

func TestLogfile(t *testing.T) {
	lf := GetAppLogfile(LEVEL_INFO)
	defer lf.Close()

	_ = lf.WriteLevelMsg("hello world for debug", LEVEL_DEBUG)
	_ = lf.WriteLevelMsg("hello world for info", LEVEL_INFO)
	_ = lf.WriteLevelMsg("hello world for warn", LEVEL_WARN)

	// 延迟1分钟
	//time.Sleep(time.Minute)
	_ = lf.WriteLevelMsg("hello world for error", LEVEL_ERROR)
	_ = lf.WriteLevelMsg("hello world for fatal", LEVEL_FATAL)

	fmt.Println(time.Now().Format("200601021504"))
}

func TestLogfile_WriteJson(t *testing.T) {
	lf := GetEventLogfile()
	defer lf.Close()

	lf2 := GetBidLogfile()
	defer lf2.Close()

	var data = struct {
		Name string
		Id   int
	}{}
	data.Name = "Jon"
	data.Id = 123

	_ = lf.WriteJson("click", data)
	_ = lf2.WriteJson("bid", data)
	_ = lf.WriteJson("click", data)
}

func BenchmarkWrite(b *testing.B) {
	lf := GetAppLogfile(LEVEL_INFO)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = lf.WriteLevelMsg("hello world", LEVEL_WARN)
		}
	})
}
