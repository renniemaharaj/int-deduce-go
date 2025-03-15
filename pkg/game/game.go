package game

import (
	"math"

	iot "github.com/renniemaharaj/int-deduce-go/pkg/iot"
)

// Game struct
type Game struct {
	Peek      int             `json:"peek"`
	Boundary  Boundary        `json:"boundary"`
	Mode      iot.Tri         `json:"stepper"`
	Query     IsMoreThanQuery `json:"query"`
	Logarithm float64         `json:"logarithm"`
	QueryText string          `json:"query_text"`
}

// Update query function
func (g *Game) updateQuery() {
	switch g.Mode {
	case iot.Boundless:
		t := Square(g.Peek)
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
	switch g.Mode {
	case iot.Spaceless:
		g.Mode = iot.Boundless

	case iot.Boundless:
		if g.Boundary.Confirmed == iot.True {
			g.Logarithm = Logarithm(g)
			g.Mode = iot.Bounded
			return
		}
	}

	// Update query
	g.updateQuery()
}

// Step up function
func (g *Game) stepUp() {
	g.Peek = int(math.Max(float64(g.Query.Term), 1))
	g.Boundary.Start = g.Query.Term
	g.Query.Set(g.Boundary.Mean(), iot.Neither)
}

// Step down function
func (g *Game) stepDown() {
	g.Boundary.Confirmed = iot.True
	g.Boundary.End = g.Query.Term
	g.Mode = iot.Bounded
}

// Step function
func (g *Game) Step() {
	switch g.Query.Confirmed {
	case iot.MoreThan:
		g.stepUp()

	case iot.LessThanOrEqual:
		g.stepDown()
	}

	// Update game state
	g.update()
}
