package command

import (
	"io"
	"log"
	"os"
	"os/exec"
)

type CommandExecuter struct {
	errLogger *log.Logger
	infLogger *log.Logger
}

func New(errLogger, infLogger *log.Logger) *CommandExecuter {
	return &CommandExecuter{
		errLogger: errLogger,
		infLogger: infLogger,
	}
}

func (me CommandExecuter) Execute(cmd *exec.Cmd, input []byte) bool {
	me.infLogger.Println("Processing message...")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		me.errLogger.Printf("Processing error: %s\n", err)
		return false
	}

	go func() {
		defer stdin.Close()
		stdin.Write(input)
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = io.MultiWriter(os.Stdout, os.Stderr)

	err = cmd.Run()

	if err != nil {
		me.errLogger.Printf("Processing error: %s\n", err)
		return false
	}

	me.infLogger.Println("Processed!")

	return true
}
