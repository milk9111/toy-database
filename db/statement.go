package db

import "fmt"

const (
	tableMaxRows uint = 4000
)

type Statement struct {
	Type        StatementType
	RowToInsert Row
}

func (statement Statement) Execute(table Table) ExecuteResult {
	switch statement.Type {
	case StatementTypeInsert:
		return statement.executeInsert(table)
	case StatementTypeSelect:
		return statement.executeSelect(table)
	}

	return ExecuteResultUnrecognizedStatement
}

func (statement Statement) executeInsert(table Table) ExecuteResult {
	if table.numRows >= tableMaxRows {
		return ExecuteResultTableFull
	}

	if _, ok := table.rows[statement.RowToInsert.ID]; !ok {
		table.rows[statement.RowToInsert.ID] = statement.RowToInsert
		table.numRows++
	} else {
		return ExecuteResultFailedInsertKeyExists
	}

	return ExecuteResultSuccess
}

func (statement Statement) executeSelect(table Table) ExecuteResult {
	fmt.Println("|ID|Username|Email|")
	for _, row := range table.rows {
		fmt.Printf("|%d|%s|%s|\n", row.ID, row.Username, row.Email)
	}

	return ExecuteResultSuccess
}
