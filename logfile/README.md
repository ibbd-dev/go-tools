# 日志保存格式

按分钟生成日志文件。

## Usages

```go
package main

import (
	"git.ibbd.net/dsp/go-tools/logfile"
)

func main() {
	lf := &logfile.Logfile{
		Filename: "/tmp/test-level.log",
		Level:    logfile.LEVEL_INFO,
	}

	lf2 := &logfile.Logfile{
		Filename: "/tmp/test.log",
	}

	_ = lf2.Write("hello world")
	_ = lf.WriteLevelMsg("hello world", logfile.LEVEL_INFO)
	_ = lf2.Write("hello world")
	_ = lf.WriteLevelMsg("hello world", logfile.LEVEL_INFO)

	lf.Close()
	lf2.Close()

	println("%+v", lf)
	println("%+v", lf2)
}
```
