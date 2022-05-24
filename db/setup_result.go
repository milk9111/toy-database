package db

import "fmt"

type SetupResultFatalError struct {
	err error
}

func newSetupResultFatalError(err error) SetupResultFatalError {
	return SetupResultFatalError{
		err: err,
	}
}

func (err SetupResultFatalError) Error() string {
	return fmt.Sprintf("Fatal error occurred during setup. %s", err.err.Error())
}
