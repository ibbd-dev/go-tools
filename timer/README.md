# 低精度定时器

## Usages

```go
import (
    "github.com/ibbd-dev/go-tools/timer"
    "time"
)

func main() {
	AddFunc(func() {
		println("--->", time.Now().Format("2006-01-02 03:04:05"))
	}, 5*time.Second)

	AddFunc(func() {
		println("===>", time.Now().Format("2006-01-02 03:04:05"))
	}, 10*time.Second)

	Start()
	Start()

	time.Sleep(time.Second * 30)
}
```

