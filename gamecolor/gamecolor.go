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
	Red
	Green
)

var smap = map[string]GColor{"Black": Black, "White": White,
	"Red": Red, "Green": Green}

// IsValid indicates the validity of a GColor value
func (c GColor) IsValid() bool {
	return c >= Black && c <= Green
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
		// TODO: Fix this to use smap in a way that doesn't involve iterating it
		return [...]string{"Black", "White", "Red", "Green"}[c]
	}
	return fmt.Sprintf("INVALID(%d)", c)
}
