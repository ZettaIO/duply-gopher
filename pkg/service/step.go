package service

import "github.com/golang/glog"

// TaskStep interface
type TaskStep interface {
	Name() string
	Run() (TaskStepResult, error)
}

// TaskStepResult is a generic result
type TaskStepResult interface {
	Marshall() []byte
}

// DummyStep for testing
type DummyStep struct{}

// Run the dummy task
func (d *DummyStep) Run() (TaskStepResult, error) {
	glog.Info("Running dummy step")
	return nil, nil
}

// Name of the step
func (d *DummyStep) Name() string {
	return "DummyStep"
}