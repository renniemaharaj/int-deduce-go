package game

import (
	iotaTypes "github.com/renniemaharaj/int-deduce-go/pkg/iotaTypes"
)

type IsMoreThanQuery struct {
	Term      int           `json:"term"`
	Confirmed iotaTypes.Tri `json:"confirmed"`
}

func (q *IsMoreThanQuery) Set(term int, c iotaTypes.Tri) {
	q.Term = term
	q.Confirmed = c
}

func (q *IsMoreThanQuery) SetConfirmed(s iotaTypes.Tri) {
	q.Confirmed = s
}
