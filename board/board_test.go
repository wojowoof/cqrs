package board

import (
	//"regexp"
	"fmt"
	"testing"
	"wojones.com/src/gamecolor"
)

func TestBoardSize(t *testing.T) {
	for c := 1; c < 15; c++ {
		for r := 1; r < 15; r++ {
			b := Board{ID: fmt.Sprintf("TestSize%dx%d", c, r),
				Columns: c, Rows: r}
			bsize := b.Squares()
			if bsize != r*c {
				t.Errorf("Incorrect size for %d x %d board: %d", c, r, bsize)
			}
		}
	}
}

func TestBoardSquareColorsAndValidity(t *testing.T) {
	const testrows = 3
	const testcols = 7
	b := Board{ID: "TestSquareColors", Columns: testcols, Rows: testrows,
		Color0: gamecolor.Black, Color1: gamecolor.White}
	for i := 0; i < testrows*testcols; i++ {

		if !b.ValidPosition(i/testcols, i%testcols) {
			t.Errorf("Position %d, %d incorrectly determined invalid",
				i/testcols, i%testcols)
		}

		idx := b.IndexOf(i/testcols, i%testcols)
		if i != idx {
			t.Errorf("Incorrect index (%d) of position %d, %d",
				idx, i/testcols, i%testcols)
		}

		sqc := b.ColorOf(i/testcols, i%testcols)
		expect := [...]gamecolor.GColor{gamecolor.Black, gamecolor.White}[i%2]
		if sqc != expect {
			t.Errorf("Incorrect square color at %d, %d (%d): %v != expected %v",
				i/testcols, i%testcols, i, sqc, expect)
		}
	}
}
