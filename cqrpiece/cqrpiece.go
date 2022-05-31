package cqrpiece

import (
	"fmt"
	"wojones.com/src/gamecolor"
)

// Cqrpiece is a badly spelled name for a structure representing
// a checkerpiece
type Cqrpiece struct {
	Name  string
	Color gamecolor.GColor
	Xp    int
	Yp    int
}

func (c *Cqrpiece) String() string {
	return fmt.Sprintf("%s (%s) @%d/%d", c.Name, c.Color, c.Xp, c.Yp)
}

// NewPiece creates a new Cqrpiece
func NewPiece(id string, color gamecolor.GColor) *Cqrpiece {
	return &Cqrpiece{Name: id, Color: color, Xp: 0, Yp: 0}
}
