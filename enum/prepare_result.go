package enum

type PrepareResult int

const (
	PrepareResultSuccess PrepareResult = iota
	PrepareResultUnrecognizedStatement
	PrepareResultSyntaxError
	PrepareResultFatalError
)
