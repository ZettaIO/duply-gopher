package gpg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ZettaIO/duply-gopher/pkg/config"
	"github.com/ZettaIO/duply-gopher/pkg/utils"
	"github.com/golang/glog"
)

const gpg_command = "gpg2"

// ImportKeys imports keys in conig
func ImportKeys(conf *config.DuplyConfig) error {
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
	err := importPublicKey(env, conf.Keys.Master.Data)
	if err != nil {
		glog.Fatalf("Failed to import master key: %v", err)
	}

	// Import host public key
	glog.Info("Importing host public key")
	err = importPublicKey(env, conf.Keys.Host.Public)
	if err != nil {
		glog.Fatalf("Failed to import host public key: %v", err)
	}

	// Import host private key
	glog.Info("Importing host private key")
	err = importPrivateKey(
		env,
		conf.Keys.Host.Private,
		conf.Keys.Host.Password)
	if err != nil {
		glog.Fatalf("Failed to import host private key: %v", err)
	}

	err = importOwnerTrusts(env)
	if err != nil {
		return err
	}
	return nil
}

func importOwnerTrusts(env []string) error {
	glog.Info("Importing owner trusts")
	cmd := exec.Command(gpg_command, "--list-keys", "--fingerprint", "--with-colons")
	cmd.Env = env
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	trusts := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "fpr") {
			s := strings.Split(line, ":")
			trusts = append(trusts, s[len(s)-2])
		}
	}

	cmd = exec.Command(gpg_command, "--import-ownertrust")
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	for _, trust := range trusts {
		io.WriteString(stdin, fmt.Sprintf("%v:6\n", trust))
	}
	stdin.Close()
	err = cmd.Wait()
	if err != nil {
		return err
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
			gpg_command,
			"--pinentry-mode=loopback",
			"--passphrase",
			passphrase,
			"--no-tty",
			"--import")
	} else {
		cmd = exec.Command(
			gpg_command,
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
func ListKeys(conf *config.DuplyConfig) {
	cmd := exec.Command(gpg_command, "--list-keys")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = conf.Env()
	cmd.Start()
	cmd.Wait()
}

// GenHostKey ...
func GenHostKey(conf *config.DuplyConfig, pw string) {
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
