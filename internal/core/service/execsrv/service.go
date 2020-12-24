package execsrv

import (
	"bytes"
	"errors"
	"time"

	"github.com/na7r1x/acectl/internal/core/domain"
	"golang.org/x/crypto/ssh"
)

type service struct {
	executor domain.Executor
}

func New() *service {
	return &service{
		executor: domain.NewExecutor(),
	}
}

func (srv *service) Exec(brk domain.Broker, cmd []string, timeout int) (string, string, error) {
	config := &ssh.ClientConfig{
		User: brk.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(brk.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(timeout) * time.Second,
	}

	// determine port, default to 22
	var _port string
	if brk.Port != "" {
		_port = brk.Port
	} else {
		_port = "22"
	}

	client, err := ssh.Dial("tcp", brk.Host+":"+_port, config)
	if err != nil {
		return "", "", errors.New("Failed to dial: " + err.Error())
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", "", errors.New("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// prep for execution
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	// exec commands

	var exitCode string
	for _, c := range cmd {
		if err := session.Run(c); err != nil {
			// return "", errors.New("Failed to run: " + err.Error())
			exitCode = err.Error()
		}
	}

	if stderr.String() != "" {
		return stdout.String(), exitCode, errors.New(stderr.String())
	}

	return stdout.String(), exitCode, nil

	// }
	// thisSession = srv.executor.Sessions[brk.Id]

}
