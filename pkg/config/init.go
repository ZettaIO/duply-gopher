package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ZettaIO/duply-gopher/pkg/gpg"
	"github.com/go-yaml/yaml"
	"github.com/golang/glog"
)

// NewConfig creats a new config
func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conf := NewDefaults()
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	err = conf.Verify()
	if err != nil {
		return nil, err
	}
	err = conf.init()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// Init ...
func (d *Config) init() error {
	path := d.Duply.ProfileRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply profile dir: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply profile dir already exist: ", path)
	}

	// Import configured public and private keys
	err := ImportKeys(&d.Duply)
	if err != nil {
		return fmt.Errorf("Failed to import keys: %v", err)
	}

	glog.Info("list Keys")
	gpg.ListKeys(d.Duply.Env())

	err = d.Duply.WriteGlobbingList()
	if err != nil {
		return fmt.Errorf("Failed to write globbing file: %v", err)
	}
	err = d.Duply.WriteProfile()
	if err != nil {
		return fmt.Errorf("Failed to write profile file: %v", err)
	}
	return nil
}

// ImportKeys imports keys in conig
func ImportKeys(conf *DuplyConfig) error {
	path := conf.GPGRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply gpg root: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply gpg root already exists: ", path)
	}

	env := conf.Env()
	// Import master public key
	glog.Info("Importing master key")
	err := gpg.ImportPublicKey(env, conf.Keys.Master.Data)
	if err != nil {
		glog.Fatalf("Failed to import master key: %v", err)
	}

	// Import host public key
	glog.Info("Importing host public key")
	err = gpg.ImportPublicKey(env, conf.Keys.Host.Public)
	if err != nil {
		glog.Fatalf("Failed to import host public key: %v", err)
	}

	// Import host private key
	glog.Info("Importing host private key")
	err = gpg.ImportPrivateKey(
		env,
		conf.Keys.Host.Private,
		conf.Keys.Host.Password)
	if err != nil {
		glog.Fatalf("Failed to import host private key: %v", err)
	}

	err = gpg.ImportOwnerTrusts(env)
	if err != nil {
		return err
	}
	return nil
}

// // GenHostKey ...
// func GenHostKey(conf *DuplyConfig, pw string) {
// 	p, err := utils.NewHavgedProcess()
// 	if err != nil {
// 		glog.Fatalf("Failed to start haveged: %v", err)
// 	}
// 	defer p.Stop()
//
// 	cmd := exec.Command("gpg", "--batch", "--gen-key")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Env = conf.Env()
// 	stdin, err := cmd.StdinPipe()
// 	if err != nil {
// 		glog.Fatalf("Failed to pass values to gpg: %v", err)
// 	}
// 	if err = cmd.Start(); err != nil {
// 		glog.Fatalf("Error executing gpg: %v", err)
// 	}
// 	io.WriteString(stdin, "Key-Type: default\n")
// 	io.WriteString(stdin, "Key-Length: default\n")
// 	io.WriteString(stdin, "Subkey-Type: default\n")
// 	io.WriteString(stdin, "Subkey-Length: default\n")
// 	io.WriteString(stdin, "Name-Real: Host Key\n")
// 	io.WriteString(stdin, "Name-Comment: Host key.\n")
// 	io.WriteString(stdin, "Name-Email: backup@host\n")
// 	io.WriteString(stdin, "Expire-Date: 0\n")
// 	io.WriteString(stdin, fmt.Sprintf("Passphrase: %s\n", pw))
// 	stdin.Close()
// 	cmd.Wait()
// }
