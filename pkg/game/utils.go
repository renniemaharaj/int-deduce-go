package game

import (
	"fmt"
	"math"
)

func Square(n int) int {
	val := math.Max(float64(n*n), 2)
	return int(val)
}

func Mean(a, b int) int {
	return (a + b) / 2
}

func Logarithm(g *Game) float64 {
	rangeSize := float64(g.Boundary.End - g.Boundary.Start)

	log := math.Log2(rangeSize) // Ensures log(0) is handled safely
	fmt.Println("Logarithm: ", log)
	return math.Max(log, 0)

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
