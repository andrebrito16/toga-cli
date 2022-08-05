package evaluator

import (
	"io"
	"os"
	"os/exec"
)

func RunWithIO(name string, args []string, in io.Reader, outw, errw io.Writer, adicionalEnv ...string) error {
	command := exec.Command(name, args...)
	command.Stdin = in
	command.Stdout = outw
	command.Stderr = errw
	if len(adicionalEnv) > 0 {
		newEnv := append(os.Environ(), adicionalEnv...)
		command.Env = newEnv
	}
	return command.Run()
}

func Run(name string, args []string, adicionalEnv ...string) error {
	return RunWithIO(name, args, os.Stdin, os.Stdout, os.Stderr, adicionalEnv...)
}
