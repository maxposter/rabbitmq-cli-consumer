package command

import (
	"os/exec"
)

type CommandFactory struct {
	Cmd  string
	Args []string
}

func (me CommandFactory) Create() *exec.Cmd {
	return exec.Command(me.Cmd, me.Args...)
}
