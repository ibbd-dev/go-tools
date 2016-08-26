package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
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
