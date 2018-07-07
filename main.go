package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
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

func preparePrompt() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v@%v#", currentUser.Username, hostname), nil
}

func main() {
	prompt, err := preparePrompt()
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occured while preparing the prompt")
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)

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
