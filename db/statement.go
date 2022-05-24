package db

import "fmt"

type StatementType int

const (
	StatementTypeInsert StatementType = iota
	StatementTypeSelect
)

type Statement struct {
	Type        StatementType
	RowToInsert Row
}

func (statement Statement) Execute(table Table) error {
	switch statement.Type {
	case StatementTypeInsert:
		return statement.executeInsert(table)
	case StatementTypeSelect:
		return statement.executeSelect(table)
	}

	return errExecuteResultUnrecognizedStatement
}

func (statement Statement) executeInsert(table Table) error {
	if table.numRows >= tableMaxRows {
		return errExecuteResultTableFull
	}

	page, err := table.rowSlot(table.numRows)
	if err != nil {
		return err
	}

	page[statement.RowToInsert.ID] = statement.RowToInsert
	table.numRows++

	return nil
}

func (statement Statement) executeSelect(table Table) error {
	fmt.Println("|ID|Username|Email|")
	for _, page := range table.pages {
		for _, row := range page {
			fmt.Printf("|%d|%s|%s|\n", row.ID, row.Username, row.Email)
		}
	}

	return nil
}
