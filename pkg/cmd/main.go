package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/steps"
	"github.com/golang/glog"
	"github.com/jasonlvhit/gocron"
	"github.com/mitchellh/multistep"
)

func main() {
	args := parseargs()
	conf, err := config.NewConfig(args.config)
	if err != nil {
		glog.Fatalf("Failed to read config file: %v", err)
	}

	service := &Service{
		Steps: []multistep.Step{
			&steps.StepDuplyBackup{Config: conf.Duply},
			&steps.StepDuplyPurge{Config: conf.Duply},
		},
		Scheduler: gocron.NewScheduler(),
	}

	service.Scheduler.Every(10).Seconds().Do(Run, service)
	service.Scheduler.Start()

	// Wait for sigterm
	glog.Info("Waiting..")
	handleSigterm()
}

type Service struct {
	Steps     []multistep.Step
	Scheduler *gocron.Scheduler
}

func Run(s *Service) {
	glog.Info("Run() ..")
	state := new(multistep.BasicStateBag)
	state.Put("Somevalue", 42)
	runner := multistep.BasicRunner{Steps: s.Steps}
	runner.Run(state)

	// Do we have a backup result in the state bag?
	if _, ok := state.GetOk("duply-backup-result"); !ok {
		glog.Info("Backup failed.. do something")
	}

	glog.Info("Run() .. DONE")
}

// ArgsParsed contains all arguments paresed from commandline
type ArgsParsed struct {
	config string
}

// Parses aruments
func parseargs() *ArgsParsed {
	args := &ArgsParsed{}
	flag.StringVar(&args.config, "config", "config.yaml", "Config file for the service (yaml)")
	flag.Parse()
	flag.Set("logtostderr", "true")
	return args
}

func handleSigterm() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	<-signalChan
	glog.Infof("Received SIGTERM, shutting down")
	os.Exit(1)
}
