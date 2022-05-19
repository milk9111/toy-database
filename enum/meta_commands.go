package enum

type MetaCommandResult int

const (
	MetaCommandResultSuccess MetaCommandResult = iota
	MetaCommandResultUnrecognizedCommand
)
