package steps

import (
	"time"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/golang/glog"
	"github.com/mitchellh/multistep"
)

type StepDuplyBackup struct {
	Config config.DuplyConfig
}

type DuplyBackupResult struct {
	StartTime                  time.Time
	EndTime                    time.Time
	ElapsedTime                time.Duration
	SourceFiles                int
	SourceFileSize             int64
	NewFiles                   int64
	NewFileSize                int64
	DeletedFiles               int64
	ChangedFiles               int64
	ChangedFileSize            int64
	ChangedDeltaSize           int64
	DeltaEntries               int64
	RawDeltaSize               int64
	TotalDestinationSizeChange int64
	Errors                     int64
}

func (s *StepDuplyBackup) Run(state multistep.StateBag) multistep.StepAction {
	glog.Info("Starting duply backup")
	// ...
	glog.Info("Env: ", s.Config.Env())
	// ...
	state.Put("duply-backup-result", &DuplyBackupResult{})
	return multistep.ActionContinue
}

func (s *StepDuplyBackup) Cleanup(multistep.StateBag) {
	// This is called after all the steps have run or if the runner is
	// cancelled so that cleanup can be performed.
}
