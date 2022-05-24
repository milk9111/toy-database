package main

import (
	"bufio"
	"fmt"
	"os"
	"toy-database/engine"
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
	engine, err := engine.Init("") //TODO: Ingest the filename if given
	if err != nil {
		panic(err)
	}

	inputBuffer := newInputBuffer()
	for {
		printPrompt()

		line, err := inputBuffer.Scan()
		if err != nil {
			panic(err)
		}

		if err = engine.ProcessInput(line); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func printPrompt() {
	fmt.Printf("db > ")
}
