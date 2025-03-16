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
		t := Square(*g.Peek)
		g.QueryText = formatQueryMoreThan(&t)
		g.Query.Set(t, iot.Neither)

	case iot.Bounded:
		if g.Boundary.Length() == 1 || g.Boundary.Spaceless() {
			if g.Query.Term == g.Boundary.Start {
				// Works in both cases due to step down side effect
				g.QueryText = formatQueryIs(&g.Boundary.End)
				g.Query.Set(g.Boundary.End, iot.Neither)
			} else {
				g.QueryText = formatQueryMoreThan(&g.Boundary.Start)
				g.Query.Set(g.Boundary.Start, iot.Neither)
			}
			break
		}
		t := g.Boundary.Mean()
		g.QueryText = formatQueryMoreThan(&t)
		g.Query.Set(t, iot.Neither)
	}
}

// Update function
func (g *Game) update() {
	switch *g.State {
	case iot.Spaceless:
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
func (g *Game) stepIn() {
	*g.Peek = g.Query.Term
	g.Boundary.Start = g.Query.Term
}

// Step down function
func (g *Game) stepOut() {
	g.Boundary.Confirmed = iot.True
	g.Boundary.End = g.Query.Term
	*g.State = iot.Bounded
}

// Step function
func (g *Game) Step() {
	switch g.Query.Confirmed {
	case iot.MoreThan:
		g.stepIn()

	case iot.LessThanOrEqual:
		g.stepOut()
	}

	// Update game state
	g.update()
}
