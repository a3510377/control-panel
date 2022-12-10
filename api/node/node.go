package node

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	STATE_BUSY = iota
	STATE_STOP
	STATE_STOPPING
	STATE_STARTING
	STATE_RUNNING
)

type Instance struct {
	ProcessArgs []string

	state          int
	Cmd            *exec.Cmd
	CommandInPipe  io.Writer
	CommandOutPipe io.ReadCloser
}

func (i *Instance) runProcess() error {
	cmd := exec.Command(i.ProcessArgs[0], i.ProcessArgs[1:]...)
	cmd.Env = os.Environ()

	i.Cmd = cmd
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	i.CommandOutPipe = stdout

	stdIn, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	i.CommandInPipe = stdIn

	return nil
}

func (i *Instance) Run() error {
	if err := i.runProcess(); err != nil {
		return err
	}

	// i.Cmd.Stdout = os.Stdout
	// i.Cmd.Stdin = os.Stdin

	go func() {
		i.Cmd.Wait()
		i.SetState(STATE_STOP)
	}()

	i.Cmd.Start()

	return nil
}

func (i *Instance) State() int {
	return i.state
}

func (i *Instance) SetState(state int) {
	if state == STATE_STOP {
		fmt.Println("\033[31mServer stopped\033[0m")
	}
	i.state = state
}
