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

// ImportKey imports a gpg key into the keychain
func ImportKey(conf *config.Config) error {
	cmd := exec.Command("gpg", "--import")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = conf.Env()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	io.WriteString(stdin, conf.Duply.Keys.Master.Data)
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
