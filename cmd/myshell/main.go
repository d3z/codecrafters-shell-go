package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var builtins = []string{"exit", "echo", "type"}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input = input[:len(input)-1] // Remove trailing newline

		command, args := parseInput(input)

		if isBuiltIn(command) {
			performBuiltin(command, args)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func parseInput(input string) (command string, args []string) {
	parsed := strings.Split(input, " ")
	return parsed[0], parsed[1:]
}

func isBuiltIn(command string) bool {
	for _, builtin := range builtins {
		if command == builtin {
			return true
		}
	}
	return false
}

func performBuiltin(command string, args []string) {
	switch command {
	case "exit":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "usage: exit <exit code>")
			return
		} else {
			exitCode, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid exit code:", args[0])
				os.Exit(1)
			}
			os.Exit(exitCode)
		}
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "type":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "usage: type <command>")
			return
		}
		if isBuiltIn(command) {
			fmt.Printf("%s is a shell builtin", command)
		} else {
			fmt.Printf("%s not found", command)
		}
	}
}
