package service

import (
	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/golang/glog"
)

// DuplyBackupStep ...
type DuplyBackupStep struct {
	config *config.DuplyConfig
}

// DuplyBackupResult ..
type DuplyBackupResult struct{}

// Duply Backup

// Run duply backup
func (s DuplyBackupStep) Run() (TaskStepResult, error) {
	glog.Info("Running duply backup")
	// ...
	return DuplyBackupResult{}, nil
}

// Name of the duply backup step
func (s DuplyBackupStep) Name() string {
	return "Duply Backup"
}

// Marshall to json
func (r DuplyBackupResult) Marshall() []byte {
	return []byte("{}")
}

// Duply Purge

// DuplyPurgeStep step
type DuplyPurgeStep struct {
	config *config.DuplyConfig
}

// DuplyPurgeResult result
type DuplyPurgeResult struct{}

// Run purge backup step
func (s DuplyPurgeStep) Run() (TaskStepResult, error) {
	glog.Info("Running duply backup")
	// ...
	return DuplyPurgeResult{}, nil
}

// Name of the purge step
func (s DuplyPurgeStep) Name() string {
	return "Duply Purge"
}

// Marshall result into json
func (r DuplyPurgeResult) Marshall() []byte {
	return []byte("{}")
}
