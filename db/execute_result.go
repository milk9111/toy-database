package db

import "fmt"

var (
	errExecuteResultUnrecognizedStatement = fmt.Errorf("Unrecognized statement reached while executing.")
	errExecuteResultTableFull             = fmt.Errorf("Table full.")
	errExecuteResultFailedInsertKeyExists = fmt.Errorf("Failed to insert - key already exists.")
	errExecuteResultPagerOutOfBounds      = fmt.Errorf("Tried to fetch page number out of bounds.")
)

type ExecuteResultFatalError struct {
	err error
}

func newExecuteResultFatalError(err error) ExecuteResultFatalError {
	return ExecuteResultFatalError{
		err: err,
	}
}

func (err ExecuteResultFatalError) Error() string {
	return fmt.Sprintf("Fatal error occurred during execution. %s", err.err.Error())
}
