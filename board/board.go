package board

import (
	"fmt"
)

type posn struct {
	column int
	row    int
}

type board struct {
	columns int
	rows    int
}

func (b board) squares() int {
	return b.columns * b.rows
}

func newBoard(id string, rows int, columns int) *board {
	return &board{rows: rows, columns: columns}
}

func init() {
	fmt.Printf("board.init()")
}
