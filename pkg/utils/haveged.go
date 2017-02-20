package utils

import (
	"os/exec"

	"github.com/golang/glog"
)

// HavgedProcess process
type HavgedProcess struct {
	done    chan error
	command *exec.Cmd
}

// Stop the process
func (h *HavgedProcess) Stop() {
	h.command.Process.Kill()
	err := <-h.done
	if err != nil {
		// Should return interrupted error
		glog.Errorf("Haveged: %v", err)
	}
}

// NewHavgedProcess runs haveged in the background
func NewHavgedProcess() (*HavgedProcess, error) {
	p := &HavgedProcess{
		done:    make(chan error),
		command: exec.Command("haveged", "-F"),
	}
	err := p.command.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		glog.Info("Haveged starting")
		p.done <- p.command.Wait()
		glog.Info("Haveged died")
	}()
	return p, nil
}
