package game

import (
	iot "github.com/renniemaharaj/int-deduce-go/pkg/iot"
)

// Game struct
type Game struct {
	Peek      *int      `json:"peek"`
	Boundary  *Boundary `json:"boundary"`
	State     *iot.Tri  `json:"stepper"`
	Query     *Query    `json:"query"`
	Logarithm *float64  `json:"logarithm"`
	QueryText string    `json:"query_text"`
}

// Update query function
func (g *Game) updateQuery() {
	switch *g.State {
	case iot.Boundless:
		QueryBoundless(g)

	case iot.Bounded:
		QueryBounded(g)
	}
}

// Update function
func (g *Game) update() {
	switch *g.State {
	case iot.Unset:
		*g.State = iot.Boundless

	case iot.Boundless:
		if g.Boundary.Confirmed == iot.True {

			*g.State = iot.Bounded
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
	*g.State = iot.Bounded
}

// Step function
func (g *Game) Step() {
	switch g.Query.Confirmed {
	case iot.Greater:
		g.up()

	case iot.LesserOr:
		g.Boundary.Confirmed = iot.True
		g.down()
	}

	// Update game state
	g.update()
}
