package db

import (
	"toy-database/db/serialize"
)

const (
	tableMaxPages = 100
	rowsPerPage   = 100
	tableMaxRows  = tableMaxPages * rowsPerPage
)

type Row struct {
	ID       uint
	Username string
	Email    string
}

type page map[uint]Row

type Table struct {
	numRows     uint
	pages       map[uint]page
	pager       serialize.Pager
	highestPage uint
}

func Open(pager serialize.Pager) (Table, error) {
	table := Table{
		numRows:     0,
		pages:       make(map[uint]page),
		pager:       pager,
		highestPage: 0,
	}

	var page page
	if err := table.pager.Read(&page); err != nil {
		return Table{}, newSetupResultFatalError(err)
	}

	if page == nil {
		page = make(map[uint]Row)
	}

	table.pages[table.highestPage] = page

	return table, nil
}

func (table Table) OverridePager(pager serialize.Pager) {
	table.pager = pager
}

func (table Table) rowSlot(rowNum uint) (page, error) {
	pageNum := rowNum / rowsPerPage
	if pageNum > tableMaxPages {
		return nil, errExecuteResultPagerOutOfBounds
	}

	// Opportunity for improvement here. To keep things simple the underlying
	// Pager is a wrapper over the 'gob' Encoder and Decoder but these can only
	// read one item at a time. We could make this use less memory unnecessarily
	// by implementing a custom reader that pulls out individual byte slices
	// from the file using the 'unsafe' package based on the actual byte length
	// of a page and an offset.
	if pageNum > table.highestPage {
		for i := uint(0); i < pageNum-table.highestPage; i++ {
			page, err := table.getNextPage()
			if err != nil {
				return nil, err
			}

			table.pages[i+table.highestPage+1] = page
		}

		table.highestPage = pageNum
	}

	return table.pages[pageNum], nil
}

func (table Table) getNextPage() (page, error) {
	var page page
	if err := table.pager.Read(&page); err != nil {
		return page, newExecuteResultFatalError(err)
	}

	if page == nil {
		page = make(map[uint]Row)
	}

	return page, nil
}
