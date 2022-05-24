package engine

import "fmt"

var (
	errPrepareResultUnrecognizedStatement = fmt.Errorf("Unrecognized keyword at start of input.")
	errPrepareResultSyntaxInvalid         = fmt.Errorf("Syntax error. Could not parse statement.")
	errPrepareResultNegativeID            = fmt.Errorf("ID must be positive.")
)

type PrepareResultFatalError struct {
	err error
}

func newPrepareResultFatalError(err error) PrepareResultFatalError {
	return PrepareResultFatalError{err: err}
}

func (err PrepareResultFatalError) Error() string {
	return fmt.Sprintf("Fatal error occurred while preparing statement. %s", err.err.Error())
}
