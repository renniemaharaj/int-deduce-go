package game

import (
	iotaTypes "github.com/renniemaharaj/int-deduce-go/pkg/iotaTypes"
)

func CreateGame() *Game {
	return &Game{}
}

type Game struct {
	Boundary  Boundary        `json:"boundary"`
	State     iotaTypes.Tri   `json:"state"`
	Query     IsMoreThanQuery `json:"query"`
	Logarithm float64         `json:"logarithm"`
	QueryText string          `json:"query_text"`
}

func (g *Game) updateQueryText() {

	g.QueryText = "Is your number more than " + PrettyPrintNumber(g.Query.Term) + "?"

	if g.Boundary.Confirmed && g.Boundary.End == g.Boundary.Start {
		g.QueryText = "Your number is " + PrettyPrintNumber(g.Boundary.Start) + "!"
	}
}

func (g *Game) update() {
	switch g.State {
	case iotaTypes.None:
		g.State = iotaTypes.Is
		g.Logarithm = Logarithm(g)
	case iotaTypes.Is:
		if g.Boundary.End-g.Boundary.Start > 1 {
			g.Logarithm = Logarithm(g)
		} else {
			g.State = iotaTypes.Not
		}
	case iotaTypes.Not:
		if g.Boundary.End-g.Boundary.Start > 1 {
			g.Logarithm = Logarithm(g)
			g.State = iotaTypes.Is
		}
	}
	g.updateQueryText()
}

func (g *Game) Step() {
	switch g.Query.Confirmed {
	case iotaTypes.Is:
		if g.Boundary.Start == g.Query.Term {
			g.Boundary.Start++
			break
		}
		g.Boundary.Start = g.Query.Term
		var nextTerm int
		if !g.Boundary.Confirmed {
			nextTerm = Square(g.Query.Term)
		} else {
			nextTerm = g.Boundary.Mean()
		}
		g.Query.Set(nextTerm, iotaTypes.None)
	case iotaTypes.Not:
		if !g.Boundary.Confirmed {
			g.Boundary.Confirmed = true
		}
		g.Boundary.End = g.Query.Term
		g.Query.Set(g.Boundary.Mean(), iotaTypes.None)
	}
	g.update()
}
