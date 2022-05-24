package engine

import (
	"fmt"
	"testing"
	"toy-database/db"
	"toy-database/mocks"

	"github.com/golang/mock/gomock"
)

type engineTestDependencies struct {
	engine engine
	pager  *mocks.MockPager
}

func setupTest(t *testing.T) engineTestDependencies {
	ctrl := gomock.NewController(t)
	pager := mocks.NewMockPager(ctrl)

	pager.EXPECT().Read(gomock.Any()).Return(nil)

	table, err := db.Open(pager)
	if err != nil {
		panic(err)
	}

	return engineTestDependencies{
		engine: engine{
			table: table,
		},
		pager: pager,
	}
}

func TestInsert10000Rows(t *testing.T) {
	dependencies := setupTest(t)

	for i := 0; i < 10000; i++ {
		if err := dependencies.engine.ProcessInput(fmt.Sprintf("insert %d test test", i)); err != nil {
			panic(err)
		}
	}

	if err := dependencies.engine.ProcessInput("select"); err != nil {
		panic(err)
	}
}

func TestInsertSinglerow(t *testing.T) {
	dependencies := setupTest(t)

	if err := dependencies.engine.ProcessInput("insert 1 test test"); err != nil {
		panic(err)
	}

	if err := dependencies.engine.ProcessInput("select"); err != nil {
		panic(err)
	}
}
