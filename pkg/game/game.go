package game

import (
	"github.com/renniemaharaj/int-deduce-go/pkg/states"
)

// Game struct
type Game struct {
	Peek      *int           `json:"peek"`
	Boundary  *Boundary      `json:"boundary"`
	State     *states.States `json:"stepper"`
	Query     *Query         `json:"query"`
	Logarithm *float64       `json:"logarithm"`
	QueryText string         `json:"query_text"`
}

// Update query function
func (g *Game) updateQuery() {
	switch *g.State {
	case states.Boundless:
		QueryBoundless(g)

	case states.Bounded:
		QueryBounded(g)
	}
}

// Update function
func (g *Game) update() {
	switch *g.State {
	case states.Unset:
		*g.State = states.Boundless

	case states.Boundless:
		if g.Boundary.Confirmed.ToBool() {

			*g.State = states.Bounded
			return
		}
	}

	*g.Logarithm = Logarithm(g)

	// Update query
	g.updateQuery()
}

// Step up function
func (g *Game) up() {
	*g.Peek = g.Query.Term
	g.Boundary.Start = g.Query.Term + 1
}

// Step down function
func (g *Game) down() {
	g.Boundary.End = g.Query.Term
	*g.State = states.Bounded
}

// Step function
func (g *Game) Step() {
	switch g.Query.Confirmed {
	case states.Greater:
		g.up()

	case states.LesserOr:
		g.Boundary.Confirmed = states.True
		g.down()
	}

	// Update game state
	g.update()
}
