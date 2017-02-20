package service

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

// Just testing
func testing() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(task1)
	s.Every(1).Second().Do(tick)

	s.Every(10).Seconds().Do(clear, s)

	fmt.Println("Starting cron")
	<-s.Start()
	fmt.Println("Stopping")
}

func task1() {
	fmt.Println("[Task1]")
}

func tick() {
	fmt.Println(".")
}

func clear(s *gocron.Scheduler) {
	s.Clear()
	s.Every(1).Second().Do(moo)
}

func moo() {
	fmt.Println("Moo")
}
