package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	AddFunc(func() {
		println("--->", time.Now().Format(time.RFC3339))
	}, 5*time.Second)

	AddFunc(func() {
		println("===>", time.Now().Format(time.RFC3339))
	}, 10*time.Second)

	AddFunc(func() {
		println("   >", time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 2)
	}, time.Second)

	time.Sleep(time.Second * 30)
}
