// see: http://studygolang.com/articles/4458
package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	doSomethings()
	go task()
	doSomethingsElse()

	select {}
}

func task() {
	i := 0
	c := cron.New()
	c.AddFunc(CRON_EVERY_FIVE_SECOND, func() {
		//c.AddFunc(CRON_02_SECOND, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
}

func doSomethings() {
	log.Println("do doSomethings ...")
}

func doSomethingsElse() {
	log.Println("do doSomethings else ...")
}
