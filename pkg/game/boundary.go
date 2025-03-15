package game

import "github.com/renniemaharaj/int-deduce-go/pkg/iot"

type Boundary struct {
	Start     int     `json:"start"`
	End       int     `json:"end"`
	Confirmed iot.Tri `json:"confirmed"`
}

func (b *Boundary) SetConfirmed(c iot.Tri) {
	b.Confirmed = c
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
