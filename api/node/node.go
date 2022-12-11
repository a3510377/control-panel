package node

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"syscall"
)

const (
	STATE_BUSY = iota
	STATE_STOP
	STATE_STOPPING
	STATE_STARTING
	STATE_RUNNING
)

type Instance struct {
	Root        string
	ProcessArgs []string

	state          int
	Cmd            *exec.Cmd
	CommandInPipe  io.Writer
	CommandOutPipe io.ReadCloser
	handles        []func(HandleEvent)
}

func New(root string, args ...string) *Instance {
	node := &Instance{
		ProcessArgs: args,
		Root:        root,
	}
	node.init()
	return node
}

func (i *Instance) init() error {
	cmd := exec.Command(i.ProcessArgs[0], i.ProcessArgs[1:]...)
	cmd.Dir = i.Root
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

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

	err := i.Cmd.Wait()
	i.SetState(STATE_STOP)
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
