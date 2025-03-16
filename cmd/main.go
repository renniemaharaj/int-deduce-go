package main

import (
	"fmt"

	"github.com/renniemaharaj/int-deduce-go/pkg/game"
	"github.com/renniemaharaj/int-deduce-go/pkg/states"
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	g := game.CreateGame() // Avoid variable shadowing

	for {
		ClearConsole()
		g.Step()
		fmt.Printf("\n--- Game Status Update ---\n")
		fmt.Printf("Query: %s\n", g.QueryText)
		fmt.Printf("Boundary: (%v, %v)\n", game.PrettyPrintNumber(g.Boundary.Start), game.PrettyPrintNumber(g.Boundary.End))
		fmt.Printf("Logarithm (Search Steps Remaining): %.2f\n", *g.Logarithm)
		fmt.Printf("Current State: %v\n", g.State)
		fmt.Println("----------------------------")
		fmt.Println("Respond with '>' (More), '<' (Less), or '/' (Exit):")

		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Input error:", err)
			continue // Prompt again instead of exiting
		}

		switch response {
		case ">":
			g.Query.SetConfirmed(states.Greater)
		case "<":
			g.Query.SetConfirmed(states.LesserOr)
		case "/":
			fmt.Println("Game exited.")
			return
		default:
			fmt.Println("Invalid response. Please enter '>' for more, '<' for less, or '/' to exit.")
			continue
		}
	}
}
