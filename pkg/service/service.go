package service

import (
	"fmt"
	"os"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/gpg"
	"github.com/golang/glog"
)

// NewService ..
func NewService(conf *config.Config) (*Service, error) {
	// Make sure config directly exists with the right permissions
	path := conf.Duply.ProfileRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply profile dir: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply profile dir already exist: ", path)
	}

	// Import configured public and private keys
	err := gpg.ImportKeys(&conf.Duply)
	if err != nil {
		return nil, fmt.Errorf("Failed to import keys: %v", err)
	}

	glog.Info("list Keys")
	gpg.ListKeys(&conf.Duply)

	err = conf.Duply.WriteGlobbingList()
	if err != nil {
		return nil, fmt.Errorf("Failed to write globbing file: %v", err)
	}
	err = conf.Duply.WriteProfile()
	if err != nil {
		return nil, fmt.Errorf("Failed to write profile file: %v", err)
	}

	s := &Service{Config: conf}
	steps := []TaskStep{
		DuplyBackupStep{},
		DuplyBackupStep{},
	}
	s.Task = NewTask(steps)
	s.Scheduler = NewTaskScheduler(s.Task)
	return s, nil
}

// Run the service
func (s *Service) Run() {
	glog.Info("Running service..")
	// Run scheduler
	s.Scheduler.RunAt("00:20")
	// Run http
	// Run MQ
}
