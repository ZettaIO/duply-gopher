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
	err := gpg.ImportKeys(conf)
	if err != nil {
		return nil, fmt.Errorf("Failed to import keys: %v", err)
	}

	glog.Info("list Keys")
	gpg.ListKeys(conf)

	err = conf.Duply.WriteGlobbingList()
	if err != nil {
		return nil, fmt.Errorf("Failed to write globbing file: %v", err)
	}

	return &Service{Config: conf}, nil
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
