package game

import "github.com/renniemaharaj/int-deduce-go/pkg/iot"

// CreateGame function
func CreateGame() *Game {
	peak := 200
	var logarithm float64 = 777
	state := iot.Spaceless

	return &Game{
		Peek:      &peak,
		State:     &state,
		Logarithm: &logarithm,
		Query:     &Query{},
		Boundary:  &Boundary{Start: 0, End: 100}, // Example default range
	}
}
