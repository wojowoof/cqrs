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

var smap = map[string]GColor{"Black": Black, "White": White}

// IsValid indicates the validity of a GColor value
func (c GColor) IsValid() bool {
	return c >= Black && c <= White
}

// FromString converts a string to a GColor
func FromString(s string) GColor {
	c, present := smap[s]
	if !present {
		return Black - 1
	}
	return c
}

func (c GColor) String() string {
	if c.IsValid() {
		return [...]string{"Black", "White"}[c]
	}
	return fmt.Sprintf("INVALID(%d)", c)
}
