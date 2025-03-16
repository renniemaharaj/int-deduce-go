package game

import "github.com/renniemaharaj/int-deduce-go/pkg/states"

type Boundary struct {
	Start     int           `json:"start"`
	End       int           `json:"end"`
	Confirmed states.States `json:"confirmed"`
}

func (b *Boundary) SetConfirmed(c *states.States) {
	b.Confirmed = *c
}

func (b *Boundary) Set(start, end int) {
	b.Start = start
	b.End = end
}

func (b *Boundary) Mean() int {
	return (b.Start + b.End) / 2
}

func (b *Boundary) Length() int {
	return b.End - b.Start
}

func (b *Boundary) Spaceless() bool {
	return b.Length() == 0
}
