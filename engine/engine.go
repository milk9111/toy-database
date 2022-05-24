package engine

import (
	"fmt"
	"os"
	"strings"
	"toy-database/db"
	"toy-database/db/serialize"
)

type Engine interface {
	ProcessInput(input string) error
}

type engine struct {
	table db.Table
}

func Init(filename string) (Engine, error) {
	pager, err := serialize.OpenPager(filename)
	if err != nil {
		return nil, err
	}

	table, err := db.Open(pager)
	if err != nil {
		return nil, err
	}

	return engine{
		table: table,
	}, nil
}

func (engine engine) ProcessInput(input string) error {
	if input[0] == '.' {
		if err := doMetaCommand(input); err != nil {
			return err
		}
	}

	statement, err := prepareStatement(input)
	if err != nil {
		return err
	}

	if err = statement.Execute(engine.table); err != nil {
		return err
	}

	return nil
}

func doMetaCommand(cmd string) error {
	if strings.Compare(cmd, ".exit") == 0 {
		os.Exit(0)
		//TODO: Close DB connection and write stored pages out to file
		return nil
	} else {
		return errMetaCommandResultUnrecognizedCommand
	}
}

func prepareStatement(cmd string) (db.Statement, error) {
	var statement db.Statement

	if strings.HasPrefix(strings.ToLower(cmd), "insert ") {
		statement.Type = db.StatementTypeInsert

		var id int
		argsAssigned, err := fmt.Sscanf(cmd, "insert %d %s %s", &id, &statement.RowToInsert.Username, &statement.RowToInsert.Email)
		if err != nil {
			return statement, newPrepareResultFatalError(err)
		}

		if argsAssigned != 3 {
			return statement, errPrepareResultSyntaxInvalid
		}

		if id < 0 {
			return statement, errPrepareResultNegativeID
		}

		statement.RowToInsert.ID = uint(id)

		return statement, nil
	}

	if strings.Compare(strings.ToLower(cmd), "select") == 0 {
		statement.Type = db.StatementTypeSelect
		return statement, nil
	}

	return statement, errPrepareResultUnrecognizedStatement
}
