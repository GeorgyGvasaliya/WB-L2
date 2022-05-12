package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		if scanner.Scan() {
			line := scanner.Text()

			commands := strings.Split(line, " | ")

			CommandsExec(commands)
		}
	}
}

func CommandsExec(commands []string) {
	for _, command := range commands {
		args := strings.Split(command, " ")
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "path required")
				continue
			}
			os.Chdir(args[1])

		case "echo":
			if len(args) < 2 {
				fmt.Fprintf(os.Stdout, "")
				continue
			}
			for i := 1; i < len(args); i++ {
				fmt.Fprintf(os.Stdout, args[i]+" ")
			}
			fmt.Println()
		case "kill":
			if len(args) < 2 {
				fmt.Fprintf(os.Stdout, "need pid")
				continue
			}
			cmd := exec.Command(args[0], args[1])

			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case `\q`:
			fmt.Fprintf(os.Stdout, "Bye!\n")
			os.Exit(1)
		default:
			cmd := exec.Command(command)

			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
