package db

type Row struct {
	ID       uint
	Username string
	Email    string
}

type Table struct {
	numRows uint
	rows    map[uint]Row
}

func NewTable() Table {
	return Table{
		numRows: 0,
		rows:    make(map[uint]Row),
	}
}
