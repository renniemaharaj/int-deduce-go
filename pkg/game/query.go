package game

import (
	iotaTypes "github.com/renniemaharaj/int-deduce-go/pkg/iot"
)

type Query struct {
	Term      int           `json:"term"`
	Confirmed iotaTypes.Tri `json:"confirmed"`
}

func (q *Query) Set(term int, c iotaTypes.Tri) {
	q.Term = term
	q.Confirmed = c
}

func (q *Query) SetConfirmed(s iotaTypes.Tri) {
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
