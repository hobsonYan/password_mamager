package utils

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var s *gocron.Scheduler

func initScheduler() {

	if s == nil {
		timezone, _ := time.LoadLocation("Asia/Shanghai")
		s = gocron.NewScheduler(timezone)
	}
}

func AddJob(job func(), corn string, tag string) {

	initScheduler()

	s.Cron(corn).Do(job)

	s.Tag(tag)

	s.StartAsync()
}

func DeleteJob(tag string) {

	initScheduler()

	err := s.RemoveByTag(tag)

	if err != nil {
		panic(err)
	} else {
		log.Println("定时任务删除成功")
	}
}
