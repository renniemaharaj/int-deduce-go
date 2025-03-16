package game

import (
	"github.com/renniemaharaj/int-deduce-go/pkg/states"
)

type Query struct {
	Term      int           `json:"term"`
	Confirmed states.States `json:"confirmed"`
}

func (q *Query) Set(term int, c states.States) {
	q.Term = term
	q.Confirmed = c
}

func (q *Query) SetConfirmed(s states.States) {
	q.Confirmed = s
}

// Format query functions
func formatQueryMoreThan(n *int) string {
	return "Is your number more than " + PrettyPrintNumber(*n) + "?"
}

// Format query functions
func formatQueryIs(n *int) string {
	return "Is your number " + PrettyPrintNumber(*n) + "?"
}
