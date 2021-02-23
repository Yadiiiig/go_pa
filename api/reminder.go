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
	// Implentation for different platforms to send 1 automatic message each day:
	// Slack, mail, discord, ...
	// It will contain a list of all the agenda items- , including the class roster of that specific day
	fmt.Println("reminding you")
}
