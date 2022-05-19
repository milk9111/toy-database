package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type inputBuffer struct {
	scanner *bufio.Scanner
}

func (buffer inputBuffer) Scan() (string, error) {
	buffer.scanner.Scan()

	return buffer.scanner.Text(), buffer.scanner.Err()
}

func newInputBuffer() inputBuffer {
	return inputBuffer{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func main() {
	inputBuffer := newInputBuffer()
	for {
		printPrompt()

		line, err := inputBuffer.Scan()
		if err != nil {
			panic(err)
		}

		if line[0] == '.' {
			switch doMetaCommand(line) {

			}
		}

		fmt.Printf("Unrecognized command '%s'.\n", line)
	}
}

func doMetaCommand(cmd string) enum.MetaCommandResult {
	if strings.Compare(cmd, ".exit") == 0 {
		os.Exit(0)
		return enum.MetaCommandResult
	} else {
		return enum.MetaCommandResultUnrecognizedCommand
	}
}

func printPrompt() {
	fmt.Printf("db > ")
}
