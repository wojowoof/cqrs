package board

import (
	"fmt"
)

type posn struct {
	column int
	row    int
}

// Board is a rectangular playing board of varibable width and height
type Board struct {
	ID      string
	Columns int
	Rows    int
}

func (b Board) String() string {
	return fmt.Sprintf("Board %s (%dx%d)", b.ID, b.Columns, b.Rows)
}

// Squares resturns the number of squares on a board
func (b Board) Squares() int {
	return b.Columns * b.Rows
}

// NewBoard returns a newly allocated board of the given size and ID
func NewBoard(id string, rows int, columns int) *Board {
	return &Board{Rows: rows, Columns: columns, ID: id}
}

func init() {
	fmt.Printf("board.init()")
}
