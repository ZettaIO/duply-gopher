package gpg

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/utils"
	"github.com/golang/glog"
)

// ImportKeys imports keys in conig
func ImportKeys(conf *config.Config) error {
	path := conf.Duply.GPGRoot()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Info("Creating duply gpg root: ", path)
		os.MkdirAll(path, 0700)
	} else {
		glog.Info("Duply gpg root already exists: ", path)
	}

	env := conf.Env()
	// Import master public key
	glog.Info("Importing master key")
	err := importPublicKey(env, conf.Duply.Keys.Master.Data)
	if err != nil {
		glog.Fatalf("Failed to import master key: %v", err)
	}

	// Import host public key
	glog.Info("Importing host public key")
	err = importPublicKey(env, conf.Duply.Keys.Host.Public)
	if err != nil {
		glog.Fatalf("Failed to import host public key: %v", err)
	}

	// Import host private key
	glog.Info("Importing host public key")
	err = importPrivateKey(
		env,
		conf.Duply.Keys.Host.Private,
		conf.Duply.Keys.Host.Password)
	if err != nil {
		glog.Fatalf("Failed to import host private key: %v", err)
	}
	return nil
}

// ImportPublicKey imports a gpg key into the keychain
func importPublicKey(env []string, keyData string) error {
	err := importKey(env, keyData, "")
	if err != nil {
		return err
	}
	return nil
}

// ImportPrivateKey imports a gpg key into the keychain
func importPrivateKey(env []string, keyData, passphrase string) error {
	err := importKey(env, keyData, passphrase)
	if err != nil {
		return err
	}
	return nil
}

func importKey(env []string, keyData, passphrase string) error {
	cmd := exec.Command("dummy")
	if passphrase != "" {
		cmd = exec.Command(
			"gpg2",
			"--pinentry-mode=loopback",
			"--passphrase",
			passphrase,
			"--no-tty",
			"--import")
	} else {
		cmd = exec.Command(
			"gpg2",
			"--pinentry-mode=loopback",
			"--no-tty",
			"--import")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	io.WriteString(stdin, keyData)
	stdin.Close()
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

// ListKeys ..
func ListKeys(conf *config.Config) {
	cmd := exec.Command("gpg", "--list-keys")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = conf.Env()
	cmd.Start()
	cmd.Wait()
}

// GenHostKey ...
func GenHostKey(conf *config.Config, pw string) {
	p, err := utils.NewHavgedProcess()
	if err != nil {
		glog.Fatalf("Failed to start haveged: %v", err)
	}
	defer p.Stop()

	cmd := exec.Command("gpg", "--batch", "--gen-key")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = conf.Env()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		glog.Fatalf("Failed to pass values to gpg: %v", err)
	}
	if err = cmd.Start(); err != nil {
		glog.Fatalf("Error executing gpg: %v", err)
	}
	io.WriteString(stdin, "Key-Type: default\n")
	io.WriteString(stdin, "Key-Length: default\n")
	io.WriteString(stdin, "Subkey-Type: default\n")
	io.WriteString(stdin, "Subkey-Length: default\n")
	io.WriteString(stdin, "Name-Real: Host Key\n")
	io.WriteString(stdin, "Name-Comment: Host key.\n")
	io.WriteString(stdin, "Name-Email: backup@host\n")
	io.WriteString(stdin, "Expire-Date: 0\n")
	io.WriteString(stdin, fmt.Sprintf("Passphrase: %s\n", pw))
	stdin.Close()
	cmd.Wait()
}
