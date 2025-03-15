package game

import (
	"math"

	iot "github.com/renniemaharaj/int-deduce-go/pkg/iot"
)

type Game struct {
	Peek      int             `json:"peek"`
	Boundary  Boundary        `json:"boundary"`
	Stepper   iot.Tri         `json:"stepper"`
	Query     IsMoreThanQuery `json:"query"`
	Logarithm float64         `json:"logarithm"`
	QueryText string          `json:"query_text"`
}

func CreateGame() *Game {
	return &Game{
		Peek:      200,
		Stepper:   iot.Spaceless,
		Logarithm: 777,
		Query:     IsMoreThanQuery{},
		Boundary:  Boundary{Start: 0, End: 100}, // Example default range
	}
}

func (g *Game) updateQuery() {
	switch g.Stepper {
	case iot.Boundless:
		t := Square(g.Peek)
		g.QueryText = "Is your number more than " + PrettyPrintNumber(t) + "?"
		g.Query.Set(t, iot.Neither)

	case iot.Bounded:
		if g.Boundary.Length() == 1 || g.Boundary.Spaceless() {
			if g.Query.Term == g.Boundary.Start {
				g.QueryText = "Is your number " + PrettyPrintNumber(g.Boundary.End) + "?"
				g.Query.Set(g.Boundary.Start, iot.Neither)
			} else {
				g.QueryText = "Is your number more than " + PrettyPrintNumber(g.Boundary.Start) + "?"
				g.Query.Set(g.Boundary.Start, iot.Neither)
			}
			break
		}
		t := g.Boundary.Mean()
		g.QueryText = "Is your number more than " + PrettyPrintNumber(t) + "?"
		g.Query.Set(t, iot.Neither)
	}
}

func (g *Game) update() {
	switch g.Stepper {
	case iot.Spaceless:
		g.Stepper = iot.Boundless

	case iot.Boundless:
		if g.Boundary.Confirmed == iot.True {
			g.Logarithm = Logarithm(g)
			g.Stepper = iot.Bounded
			return // Prevent infinite recursion
		}

	case iot.Bounded:
		if g.Boundary.Spaceless() || g.Boundary.Length() == 1 {
			if g.Peek == g.Boundary.End {
				g.Stepper = iot.Spaceless
				g.Logarithm = 0
			}
		} else {
			g.Logarithm = Logarithm(g)
		}
	}
	g.updateQuery()
}

func (g *Game) Step() {
	switch g.Query.Confirmed {
	case iot.MoreThan:
		g.Peek = int(math.Max(float64(g.Peek+1), 1))
		g.Boundary.Start = g.Query.Term
		g.Query.Set(g.Boundary.Mean(), iot.Neither)

	case iot.LessThanOrEqual:
		g.Boundary.Confirmed = iot.True
		g.Boundary.End = g.Query.Term
		g.Stepper = iot.Bounded
	}
	g.update()
}
