package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func runReminders() {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Day().At("22:29").Do(doSomething)
	s.StartAsync()
	fmt.Println(job.ScheduledAtTime())
}

func doSomething() {
	fmt.Println("reminding you")
}
