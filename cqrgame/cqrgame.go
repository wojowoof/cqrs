package cqrgame

import (
	"fmt"
)

// Cqrgame - a game of checkers
type Cqrgame struct {
	ID string
}

func (g *Cqrgame) String() string {
	return fmt.Sprintf("Game %s", g.ID)
}
