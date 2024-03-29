package cmd

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
	"gopub3.0/mlog"

	"io/ioutil"
	"os/exec"
	"strings"
)

// run command with local
func RunLocal(command string) (output string, err error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	mlog.Flog("localCommand", "[local command run]", command)

	if err = cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return "", err
	}

	stdin.Close()

	outBytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	errBytes, _ := ioutil.ReadAll(stderr)
	stderr.Close()

	if err = cmd.Wait(); err != nil {
		mlog.Flog("localCommand", "[local command result]", strings.TrimSpace(string(errBytes)))
		return "", errors.New(strings.TrimSpace(string(errBytes)))
	}
	mlog.Flog("localCommand", "[local command result]", string(outBytes))

	// fmt.Println("Execute finished:" + string(outBytes))
	return string(outBytes), nil
}

// run remote command
func RunRemote(session *ssh.Session, command string) (result string, err error) {
	// stdin, _ := session.StdinPipe()
	stdout, _ := session.StdoutPipe()
	stderr, _ := session.StderrPipe()
	mlog.Flog("remoteCommand", "[remote command run]", command)

	if err = session.Run(command); err != nil {
		errBytes, _ := ioutil.ReadAll(stderr)
		mlog.Flog("remoteCommand", "[remote command result]", string(errBytes))
		return err.Error() + ":" + string(errBytes), err
	}

	// stdin.Close()

	outBytes, _ := ioutil.ReadAll(stdout)
	mlog.Flog("remoteCommand", "[remote command result]", string(outBytes))

	return string(outBytes), nil
}
