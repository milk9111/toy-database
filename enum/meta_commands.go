package enum

type MetaCommandResult int

const (
	MetaCommandResultSuccess             MetaCommandResult = 1
	MetaCommandResultUnrecognizedCommand MetaCommandResult = 2
)
