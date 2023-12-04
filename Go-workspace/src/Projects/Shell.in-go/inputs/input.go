package inputs

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecInput(input string) error {

	// Remove the newline character

	input = strings.TrimSuffix(input, "\n")

	// Split the input to seperate the command and the arguments

	args := strings.Split(input, " ")

	// Check for built-in commands

	switch args[0] {

	case "cd":
		// if input is cd --- change to home directory if with empty path not supported yet.

		if len(args) < 2 {
			return errors.New("path required")
		}

		// change the directory and return error

		return os.Chdir(args[1])

	case "exit":

		os.Exit(0)
	}

	// Prepare the command to execute

	cmd := exec.Command(args[0], args[1:]...)

	// set the correct output device

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
