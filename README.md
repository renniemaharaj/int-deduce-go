# Int Deduce

A number guessing game implementation in Go that uses binary search and logarithmic scaling to efficiently guess a number that a player is thinking of.

## Overview

Package game implements a number guessing game using binary search and logarithmic scaling.

The Game struct represents the main game state with the following fields:
- `Peek`: Current number being evaluated (pointer to int)
- `Boundary`: Range limits for the search (pointer to Boundary)
- `State`: Current game state using iot.Tri (Spaceless/Boundless/Bounded)
- `Query`: Handles the current question state (pointer to Query)
- `Logarithm`: Tracks search progress (pointer to float64)
- `QueryText`: The current question for the player (string)

## Game States

- **Spaceless**: Initial state when the game starts
- **Boundless**: Exploring the number space to establish boundaries
- **Bounded**: Operating within known boundaries using binary search

## Core Functions

- `updateQuery()`: Updates game's question based on current state and boundary conditions
- `update()`: Manages state transitions and logarithmic progression
- `up()`: Updates game state when number is higher than current guess
- `down()`: Updates game state when number is lower than or equal to current guess
- `Step()`: Main game progression function handling player responses

## Installation

```bash
go get github.com/renniemaharaj/int-deduce-go
```

## Usage

```go
import "github.com/renniemaharaj/int-deduce-go/pkg/game"

func main() {
    g := game.CreateGame()
    g.Step() // Advances game state
    // Process g.QueryText for player input
}
```

## Features

- Binary search with logarithmic scaling
- Dynamic boundary adjustment
- JSON serialization support
- State management using iot.Tri type
- Smart question formatting
- Efficient number space exploration

