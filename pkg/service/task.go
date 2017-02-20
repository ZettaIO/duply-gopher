package service

// Task that can be scheduled
type Task struct {
	Results []TaskStepResult
	Steps   []TaskStep
}

// NewTask with steps
func NewTask(steps []TaskStep) Task {
	return Task{
		Steps:   steps,
		Results: make([]TaskStepResult, len(steps)),
	}
}

// Run the task
func (t *Task) Run() error {
	for i, step := range t.Steps {
		result, err := step.Run()
		if err != nil {
			return err
		}
		t.Results[i] = result
	}
	return nil
}
