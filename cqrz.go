package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	// board "wojones.com/src/board"
	"wojones.com/src/cqrpiece"
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

	//nb := board.newBoard(8, 8)
	np := cqrpiece.newPiece("Black0")
	fmt.Printf("Piece: %v", np)
}

func init() {
	progname = filepath.Base(os.Args[0])
	flag.BoolVar(&quiet, "q", false, "Be quiet!")
	flag.Parse()
}
