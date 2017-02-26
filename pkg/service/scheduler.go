package service

import (
	"github.com/golang/glog"
	"github.com/jasonlvhit/gocron"
)

type TaskScheduler struct {
	scheduler *gocron.Scheduler
	task      *Task
}

func NewTaskScheduler(task *Task) TaskScheduler {
	return TaskScheduler{
		task:      task,
		scheduler: gocron.NewScheduler(),
	}
}

func (t TaskScheduler) RunAt(timestamp string) {
	glog.Info("Updating scheduler: ", timestamp)
	t.scheduler.Every(10).Seconds().Do(func() {
		glog.Info("Tick..")
		t.task.Run()
	})
	t.scheduler.Start()
}
