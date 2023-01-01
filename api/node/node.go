package node

import (
	"bufio"
	"io"
	"os"
	"os/exec"

	"github.com/a3510377/control-panel/service/id"
)

const (
	STATE_BUSY = iota
	STATE_STOP
	STATE_STOPPING
	STATE_STARTING
	STATE_RUNNING
)

type Instance struct {
	ID          id.ID
	Name        string
	Root        string
	ProcessArgs []string

	state          int
	Cmd            *exec.Cmd
	CommandInPipe  io.Writer
	CommandOutPipe io.ReadCloser
	handles        []func(HandleEvent)
}

func (i *Instance) init() error {
	cmd := exec.Command(i.ProcessArgs[0], i.ProcessArgs[1:]...)
	cmd.Dir = i.Root
	cmd.Env = os.Environ()
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // TODO fix linux build error

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
	i.SetState(STATE_STARTING)
	if err := i.init(); err != nil {
		return err
	}

	go func() {
		// FIXME this should be done in a separate goroutine
		buf := bufio.NewScanner(i.CommandOutPipe)
		for buf.Scan() {
			i.Dispatch(HandleEvent{Name: MessageEvent, Data: buf.Text()}) // call `MessageEvent`
		}
	}()

	if err := i.Cmd.Start(); err != nil {
		i.SetState(STATE_STOP)
		return err
	} else {
		i.SetState(STATE_RUNNING)
	}

	defer i.Kill()

	err := i.Cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (i *Instance) Stop() error {
	i.SetState(STATE_STOPPING)
	return i.Cmd.Process.Signal(os.Interrupt)
}

func (i *Instance) Kill() error {
	i.SetState(STATE_STOP)
	return i.Cmd.Process.Kill()
}

func (i *Instance) State() int {
	return i.state
}

func (i *Instance) SetState(state int) {
	i.state = state
	i.Dispatch(HandleEvent{Name: StateEvent, Data: i.state}) // call `StateEvent`
}
