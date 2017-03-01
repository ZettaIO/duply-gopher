package gpg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/glog"
)

const gpg_command = "gpg2"

func ImportOwnerTrusts(env []string) error {
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
func ImportPublicKey(env []string, keyData string) error {
	err := ImportKey(env, keyData, "")
	if err != nil {
		return err
	}
	return nil
}

// ImportPrivateKey imports a gpg key into the keychain
func ImportPrivateKey(env []string, keyData, passphrase string) error {
	err := ImportKey(env, keyData, passphrase)
	if err != nil {
		return err
	}
	return nil
}

func ImportKey(env []string, keyData, passphrase string) error {
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
func ListKeys(env []string) {
	cmd := exec.Command(gpg_command, "--list-keys")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	cmd.Start()
	cmd.Wait()
}
