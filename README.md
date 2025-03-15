# Int Deduce

A number guessing game implementation in Go that uses binary search and logarithmic scaling to efficiently guess a number that a player is thinking of.

## Overview

Package game implements a number guessing game using binary search and logarithmic scaling.

The Game struct represents the main game state with the following fields:
- `Peek`: Current number being evaluated
- `Boundary`: Range limits for the search
- `Mode`: Current game state (Spaceless/Boundless/Bounded)
- `Query`: Handles the current question state
- `Logarithm`: Tracks search progress
- `QueryText`: The current question for the player

## Game Modes

- **Spaceless**: Initial state when the game starts
- **Boundless**: Exploring the number space to establish boundaries
- **Bounded**: Operating within known boundaries

## Core Functions

- `updateQuery()`: Updates the game's current question based on game mode
- `update()`: Updates game state and progression
- `stepUp()`: Handles progression when number is higher than current guess
- `stepDown()`: Handles progression when number is lower than or equal to current guess
- `Step()`: Main game progression function that processes player responses

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

- Binary search algorithm with logarithmic scaling
- Dynamic boundary adjustment
- JSON serialization support
- Pretty number formatting
- Efficient number guessing strategy
