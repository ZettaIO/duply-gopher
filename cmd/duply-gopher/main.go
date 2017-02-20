package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/service"
	"github.com/golang/glog"
)

func main() {
	args := parseargs()
	conf, err := config.NewConfig(args.config)
	if err != nil {
		glog.Fatalf("Failed to read config file: %v", err)
	}

	srv := service.NewService(conf)
	// Configure Tasks and steps
	srv.Run()

	// Wait forever until SIGTERM
	glog.Info("All done. Waiting..")
	// handleSigterm()
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
