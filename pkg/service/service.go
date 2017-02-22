package service

import (
	"os"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/gpg"
	"github.com/golang/glog"
)

// NewService ..
func NewService(conf *config.Config) *Service {
	// Make sure config directly exists with the right permissions
	path := conf.Duply.GPGRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply gpg root: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply gpg root already exists: ", path)
	}
	path = conf.Duply.ProfileRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply profile dir: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply profile dir already exist: ", path)
	}

	// Import master public key
	glog.Info("Importing master key")
	err := gpg.ImportKey(conf)
	if err != nil {
		glog.Fatalf("Failed to import master key: %v", err)
	}

	glog.Info("Generate host key")
	gpg.GenHostKey(conf, "bolle")

	glog.Info("list Keys")
	gpg.ListKeys(conf)

	err = conf.Duply.WriteGlobbingList()
	if err != nil {
		glog.Fatal("Failed to write globbing file: ", err)
	}

	return &Service{Config: conf}
}

// Service ..
type Service struct {
	Config *config.Config
}

// Run the service
func (s *Service) Run() {
	// Run scheduler
	// Run MQ
	// Run http

}
