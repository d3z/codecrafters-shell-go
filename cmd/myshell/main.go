package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		
		if command == "exit" {
			if len(args) != 1 {
				fmt.Fprintln(os.Stderr, "usage: exit <exit code>")
				continue
			} else {
				exitCode, err := strconv.Atoi(args[0])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Invalid exit code:", args[0])
					os.Exit(1)
				}

				os.Exit(exitCode)
			}
		}

		fmt.Printf("%s: command not found\n", input)
	}
}

func parseInput(input string) (command string, args []string) {
	parsed := strings.Split(input, " ")
	return parsed[0], parsed[1:]
}