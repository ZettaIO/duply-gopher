package main

import (
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/steps"
	"github.com/golang/glog"
	"github.com/mitchellh/multistep"
	"github.com/robfig/cron"
)

func main() {
	args := parseargs()
	conf, err := config.NewConfig(args.config)
	if err != nil {
		glog.Fatalf("Failed to read config file: %v", err)
	}

	job := &BackupJob{
		Steps: []multistep.Step{
			&steps.StepDuplyBackup{Config: conf.Duply},
			&steps.StepDuplyPurge{Config: conf.Duply},
		},
		Scheduler: cron.New(),
	}
	job.Scheduler.AddFunc(conf.Duply.RunAt, func() { run(job) })
	job.Scheduler.Start()

	// Wait for sigterm
	glog.Info("Waiting..")
	handleSigterm()
}

type BackupJob struct {
	Steps     []multistep.Step
	Scheduler *cron.Cron
}

func run(job *BackupJob) error {
	glog.Info("---[ Run() Start ]---")
	state := new(multistep.BasicStateBag)
	state.Put("Somevalue", 42)
	runner := multistep.BasicRunner{Steps: job.Steps}
	runner.Run(state)

	// Do we have any errors in the bag?
	if err, ok := state.GetOk("error"); ok {
		return err.(error)
	}

	// Do we have a backup result in the state bag?
	r, ok := state.GetOk("duply-backup-result")
	if !ok {
		glog.Info("Backup failed.. do something")
	}

	// Print out backup result as json
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	glog.Info("BackupResult: ", string(b))

	glog.Info("---[ Run() End ]---")
	return nil
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
