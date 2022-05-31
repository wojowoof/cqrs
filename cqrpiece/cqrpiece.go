package cqrpiece

import (
	"fmt"
)

// Cqrpiece is a badly spelled name for a structure representing
// a checkerpiece
type Cqrpiece struct {
	Name string
	Xp   int
	Yp   int
}

func (c *Cqrpiece) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Xp, c.Yp)
}

// NewPiece creates a new Cqrpiece
func NewPiece(id string) *Cqrpiece {
	return &Cqrpiece{Name: id, Xp: 0, Yp: 0}
}
