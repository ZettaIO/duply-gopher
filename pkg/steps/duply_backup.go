package steps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/golang/glog"
	"github.com/mitchellh/multistep"
)

type StepDuplyBackup struct {
	Config config.DuplyConfig
}

type DuplyBackupResult struct {
	StartTime                  time.Time `json:"start_time"`
	EndTime                    time.Time `json:"end_time"`
	ElapsedTime                float64   `json:"elapsed_time"`
	SourceFiles                int64     `json:"source_files"`
	SourceFileSize             int64     `json:"source_file_size"`
	NewFiles                   int64     `json:"new_files"`
	NewFileSize                int64     `json:"new_file_size"`
	DeletedFiles               int64     `json:"deleted_files"`
	ChangedFiles               int64     `json:"changed_files"`
	ChangedFileSize            int64     `json:"changed_file_size"`
	ChangedDeltaSize           int64     `json:"changed_delta_size"`
	DeltaEntries               int64     `json:"delta_entries"`
	RawDeltaSize               int64     `json:"raw_delta_size"`
	TotalDestinationSizeChange int64     `json:"toal_destination_size_change"`
	Errors                     int64     `json:"errors"`
}

func (s *StepDuplyBackup) Run(state multistep.StateBag) multistep.StepAction {
	glog.Info("Starting duply backup")
	// ...
	// r, err := parseOutput(data)
	// ...
	state.Put("duply-backup-result", &DuplyBackupResult{})
	return multistep.ActionContinue
}

func (s *StepDuplyBackup) Cleanup(multistep.StateBag) {
	// This is called after all the steps have run or if the runner is
	// cancelled so that cleanup can be performed.
}

func parseOutput(data string) (*DuplyBackupResult, error) {
	r := &DuplyBackupResult{}
	for i, line := range strings.Split(data, "\n") {
		s := strings.Split(strings.TrimSpace(line), " ")
		fmt.Println(i, " ", s)
		if len(s) < 2 {
			continue
		}
		switch s[0] {
		case "StartTime":
			t, _ := strconv.ParseFloat(s[1], 64)
			r.StartTime = time.Unix(int64(t), 0)
		case "EndTime":
			t, _ := strconv.ParseFloat(s[1], 64)
			r.EndTime = time.Unix(int64(t), 0)
		case "ElapsedTime":
			r.ElapsedTime, _ = strconv.ParseFloat(s[1], 64)
		case "SourceFiles":
			r.SourceFiles, _ = strconv.ParseInt(s[1], 10, 64)
		case "SourceFileSize":
			r.SourceFileSize, _ = strconv.ParseInt(s[1], 10, 64)
		case "NewFiles":
			r.NewFiles, _ = strconv.ParseInt(s[1], 10, 64)
		case "NewFileSize":
			r.NewFileSize, _ = strconv.ParseInt(s[1], 10, 64)
		case "DeletedFiles":
			r.DeletedFiles, _ = strconv.ParseInt(s[1], 10, 64)
		case "ChangedFiles":
			r.ChangedFiles, _ = strconv.ParseInt(s[1], 10, 64)
		case "ChangedFileSize":
			r.ChangedFileSize, _ = strconv.ParseInt(s[1], 10, 64)
		case "ChangedDeltaSize":
			r.ChangedDeltaSize, _ = strconv.ParseInt(s[1], 10, 64)
		case "DeltaEntries":
			r.DeltaEntries, _ = strconv.ParseInt(s[1], 10, 64)
		case "RawDeltaSize":
			r.RawDeltaSize, _ = strconv.ParseInt(s[1], 10, 64)
		case "TotalDestinationSizeChange":
			r.TotalDestinationSizeChange, _ = strconv.ParseInt(s[1], 10, 64)
		case "Errors":
			r.Errors, _ = strconv.ParseInt(s[1], 10, 64)
		}
	}
	return r, nil
}
