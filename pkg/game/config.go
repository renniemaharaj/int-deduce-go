package game

import "github.com/renniemaharaj/int-deduce-go/pkg/iot"

// CreateGame function
func CreateGame() *Game {
	return &Game{
		Peek:      200,
		Mode:      iot.Spaceless,
		Logarithm: 777,
		Query:     IsMoreThanQuery{},
		Boundary:  Boundary{Start: 0, End: 100}, // Example default range
	}
}

// Format query functions
func formatQueryMoreThan(n *int) string {
	return "Is your number more than " + PrettyPrintNumber(*n) + "?"
}

// Format query functions
func formatQueryIs(n *int) string {
	return "Is your number " + PrettyPrintNumber(*n) + "?"
}
