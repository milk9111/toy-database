package db

type StatementType int

const (
	StatementTypeInsert StatementType = iota
	StatementTypeSelect
)

type ExecuteResult int

const (
	ExecuteResultSuccess ExecuteResult = iota
	ExecuteResultTableFull
	ExecuteResultUnrecognizedStatement
	ExecuteResultFailedInsertKeyExists
)
