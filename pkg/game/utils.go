package game

import (
	"fmt"
	"math"

	"github.com/renniemaharaj/int-deduce-go/pkg/states"
)

func Square(n int) int {
	val := n * n
	return int(val)
}

func Mean(a, b int) int {
	return (a + b) / 2
}

func Logarithm(g *Game) float64 {
	rangeSize := float64(g.Boundary.End - g.Boundary.Start)

	log := math.Log2(rangeSize) // Ensures log(0) is handled safely
	return math.Max(log, 0)

}

// Query boundless function
func QueryBoundless(g *Game) {
	t := Square(*g.Peek)
	g.QueryText = formatQueryMoreThan(&t)
	g.Query.Set(t, states.Unset)
}

// Query bounded function
func QueryBounded(g *Game) {
	if g.Boundary.Length() == 1 {
		g.QueryText = formatQueryMoreThan(&g.Boundary.Start)
		g.Query.Set(g.Boundary.Start, states.Unset)
		return
	}

	if g.Boundary.Spaceless() {
		g.QueryText = formatQueryIs(&g.Boundary.Start)
		g.Query.Set(g.Boundary.Start, states.Unset)
		return
	}

	t := g.Boundary.Mean()
	g.QueryText = formatQueryMoreThan(&t)
	g.Query.Set(t, states.Unset)
}

func PrettyPrintNumber(n int) string {
	// Convert to string first
	numStr := fmt.Sprintf("%d", n)

	// Handle numbers less than 1000
	if len(numStr) <= 3 {
		return string(numStr)
	}

	// Add commas for larger numbers
	result := ""
	for i, digit := range numStr {
		if i > 0 && (len(numStr)-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}

	return result
}
