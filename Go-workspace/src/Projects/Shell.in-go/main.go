package main

import (
	"bufio"
	"fmt"
	"os"
	"projects/Shell.in-go/inputs"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("% ")

		// read from the keyboard

		input, err := reader.ReadString('\n')

		if err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input

		if err = inputs.ExecInput(input); err != nil {

			fmt.Fprintln(os.Stderr, err)
		}

	}
}
