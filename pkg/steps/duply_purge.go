package steps

import (
	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/golang/glog"
	"github.com/mitchellh/multistep"
)

type StepDuplyPurge struct {
	Config config.DuplyConfig
}

func (s *StepDuplyPurge) Run(state multistep.StateBag) multistep.StepAction {
	glog.Info("Starting duply purge")
	// ...
	return multistep.ActionContinue
}

func (s *StepDuplyPurge) Cleanup(multistep.StateBag) {
	// This is called after all the steps have run or if the runner is
	// cancelled so that cleanup can be performed.
}
