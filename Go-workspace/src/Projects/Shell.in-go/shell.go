package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("% ")
		
		// read from the keyboard

		input , err := reader.ReadString('\n')

		if err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input
		
		if err = execInput(input); err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

		

	}
}


func execInput(input string) error {

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

	cmd := exec.Command(args[0], args[1:]... )

	// set the correct output device

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func helloWorld(x string) string {

	if x == "hello" {
		return "Hi"
	}
        
}
