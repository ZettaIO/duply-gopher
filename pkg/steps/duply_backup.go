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
	StartTime                  time.Time     `json:"start_time"`
	EndTime                    time.Time     `json:"end_time"`
	ElapsedTime                time.Duration `json:"elapsed_time"`
	SourceFiles                int           `json:"source_files"`
	SourceFileSize             int64         `json:"source_file_size"`
	NewFiles                   int64         `json:"new_files"`
	NewFileSize                int64         `json:"new_file_size"`
	DeletedFiles               int64         `json:"deleted_files"`
	ChangedFiles               int64         `json:"changed_files"`
	ChangedFileSize            int64         `json:"changed_file_size"`
	ChangedDeltaSize           int64         `json:"changed_delta_size"`
	DeltaEntries               int64         `json:"delta_entries"`
	RawDeltaSize               int64         `json:"raw_delta_size"`
	TotalDestinationSizeChange int64         `json:"toal_destination_size_change"`
	Errors                     int64         `json:"errors"`
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
