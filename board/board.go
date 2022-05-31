package board

import (
	"fmt"
	"strings"
	"wojones.com/src/gamecolor"
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
	Color0  gamecolor.GColor
	Color1  gamecolor.GColor
}

func (b Board) String() string {
	return fmt.Sprintf("Board %s (%dx%d)", b.ID, b.Columns, b.Rows)
}

// Dump returns an ASCII-style string representation of a board
func (b Board) Dump() string {
	s := "+" + strings.Repeat("-+", b.Columns) + "\n"
	for i := 0; i < b.Rows; i++ {
		s += "|" + strings.Repeat(" |", b.Columns) + "\n"
		s += "+" + strings.Repeat("-+", b.Columns) + "\n"
	}
	return s
}

// ValidPosition indicates if the given position is on the board
func (b Board) ValidPosition(r int, c int) bool {
	return c >= 0 && c < b.Columns && r >= 0 && r < b.Rows
}

// Squares resturns the number of squares on a board
func (b Board) Squares() int {
	return b.Columns * b.Rows
}

// ColorOf returns the color of the square at the given column/row
func (b Board) ColorOf(row int, col int) gamecolor.GColor {
	if ((row+1)*b.Columns+col+1)%2 == 0 {
		return b.Color0
	}
	return b.Color1
}

// IndexOf return the index of the given column/row
func (b Board) IndexOf(row int, col int) int {
	return (row * b.Columns) + col
}

// NewBoard returns a newly allocated board of the given size and ID
func NewBoard(id string, rows int, columns int) *Board {
	return &Board{Rows: rows, Columns: columns, ID: id,
		Color0: gamecolor.Black, Color1: gamecolor.White}
}

func init() {
	// fmt.Printf("board.init()\n")
}
