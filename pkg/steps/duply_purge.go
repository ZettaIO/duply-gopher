package steps

import (
	"fmt"
	"os/exec"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/golang/glog"
	"github.com/mitchellh/multistep"
)

type StepDuplyPurge struct {
	Config config.DuplyConfig
}

func (s *StepDuplyPurge) Run(state multistep.StateBag) multistep.StepAction {
	glog.Info("Starting duply purge")
	cmd := exec.Command("duply", "profile", "purge", "--force")
	cmd.Env = s.Config.Env()
	err := cmd.Run()
	if err != nil {
		state.Put("error", fmt.Errorf("Failed to run duply purge: ", err))
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepDuplyPurge) Cleanup(state multistep.StateBag) {
	// This is called after all the steps have run or if the runner is
	// cancelled so that cleanup can be performed.
	glog.Info("Cleanup: duply backup")
}
