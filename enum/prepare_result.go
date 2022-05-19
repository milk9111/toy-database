package enum

type PrepareResult int

const (
	PrepareResultSuccess               PrepareResult = 1
	PrepareResultUnrecognizedStatement PrepareResult = 2
)
