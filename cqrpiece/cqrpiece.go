package cqrpiece

import (
	"fmt"
)

type cqrpiece struct {
	name string
	xp   int
	yp   int
}

func (c *cqrpiece) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.name, c.xp, c.yp)
}

func newPiece(id string) *cqrpiece {
	return &cqrpiece{name: id, xp: 0, yp: 0}
}
