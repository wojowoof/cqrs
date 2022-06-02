package cqrgame

import (
	"fmt"
)

// Cqrgame - a game of checkers
type Cqrgame struct {
	ID      string
	Player0 string
	Player1 string
}

func (g *Cqrgame) String() string {
	return fmt.Sprintf("Game %s between %s and %s", g.ID,
		g.Player0, g.Player1)
}
