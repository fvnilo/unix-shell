package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	cmdParts := strings.Split(input, " ")

	switch cmdParts[0] {
	case "cd":
		// TODO: handle empty param as cd to home
		if len(cmdParts) < 2 {
			return errors.New("path required")
		}
		err := os.Chdir(cmdParts[1])
		if err != nil {
			return err
		}

		return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		err = execCommand(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
