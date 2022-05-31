package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"wojones.com/src/board"
	"wojones.com/src/cqrgame"
	"wojones.com/src/cqrpiece"
	"wojones.com/src/gamecolor"
)

var (
	progname = "unknown"
	quiet    = false
)

func main() {

	fmt.Printf("%s: let's play%s!\n", progname, func() string {
		if quiet {
			return "(quietly)"
		}
		return ""
	}())

	np := cqrpiece.NewPiece("Black0", gamecolor.Black)
	fmt.Printf("Piece: %v\n", np)

	//cboard := board.NewBoard("321", 8, 10)
	cboard := board.Board{ID: "CqrBoard", Columns: 6, Rows: 4}
	fmt.Printf("Board: %v has %v squares\n", cboard, cboard.Squares())
	fmt.Printf("Board:\n")
	fmt.Printf("%s\n", cboard.Dump())

	gamep := &cqrgame.Cqrgame{ID: "game1"}
	fmt.Printf("New game: %v\n", gamep)
}

func init() {
	progname = filepath.Base(os.Args[0])
	flag.BoolVar(&quiet, "q", false, "Be quiet!")
	flag.Parse()
}
