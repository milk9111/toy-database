package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"toy-database/db"
	"toy-database/enum"
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
	table := db.NewTable()

	inputBuffer := newInputBuffer()
	for {
		printPrompt()

		line, err := inputBuffer.Scan()
		if err != nil {
			panic(err)
		}

		if line[0] == '.' {
			switch doMetaCommand(line) {
			case enum.MetaCommandResultSuccess:
				continue
			case enum.MetaCommandResultUnrecognizedCommand:
				fmt.Printf("Unrecognized command '%s' \n", line)
				continue
			}
		}

		var statement db.Statement
		switch prepareStatement(line, &statement) {
		case enum.PrepareResultSuccess:
			break
		case enum.PrepareResultUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", line)
			continue
		case enum.PrepareResultSyntaxError:
			fmt.Println("Syntax error. Could not parse statement.")
			continue
		case enum.PrepareResultFatalError:
			fmt.Printf("Fatal error occurred while preparing statement '%s'.\n", line)
			continue
		}

		switch statement.Execute(table) {
		case db.ExecuteResultSuccess:
			fmt.Println("Executed.")
		case db.ExecuteResultUnrecognizedStatement:
			fmt.Println("Error: Unrecognized statement reached while executing.")
		case db.ExecuteResultTableFull:
			fmt.Println("Error: Table full.")
		case db.ExecuteResultFailedInsertKeyExists:
			fmt.Println("Error: Failed to insert - key already exists.")
		}
	}
}

func printPrompt() {
	fmt.Printf("db > ")
}

func doMetaCommand(cmd string) enum.MetaCommandResult {
	if strings.Compare(cmd, ".exit") == 0 {
		os.Exit(0)
		return enum.MetaCommandResultSuccess
	} else {
		return enum.MetaCommandResultUnrecognizedCommand
	}
}

func prepareStatement(cmd string, statement *db.Statement) enum.PrepareResult {
	if strings.HasPrefix(strings.ToLower(cmd), "insert ") {
		statement.Type = db.StatementTypeInsert
		argsAssigned, err := fmt.Sscanf(cmd, "insert %d %s %s", &statement.RowToInsert.ID, &statement.RowToInsert.Username, &statement.RowToInsert.Email)
		if err != nil {
			fmt.Println(err.Error())
			return enum.PrepareResultFatalError
		}

		if argsAssigned != 3 {
			return enum.PrepareResultSyntaxError
		}

		return enum.PrepareResultSuccess
	}

	if strings.Compare(strings.ToLower(cmd), "select") == 0 {
		statement.Type = db.StatementTypeSelect
		return enum.PrepareResultSuccess
	}

	return enum.PrepareResultUnrecognizedStatement
}
