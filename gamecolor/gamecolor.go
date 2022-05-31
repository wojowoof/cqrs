package gamecolor

import (
	"fmt"
)

// GColor is a color for a game; a player and his pieces are a specific
// "color", allowing them to be distinguished from other colors
type GColor int

// Colors for a player/pieces: Black, White
const (
	Black = iota
	White
)

// IsValid indicates the validity of a GColor value
func (c GColor) IsValid() bool {
	return c >= Black && c <= White
}

func (c GColor) String() string {
	if c.IsValid() {
		return [...]string{"Black", "White"}[c]
	}
	return fmt.Sprintf("INVALID(%d)", c)
}
