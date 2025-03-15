Package game implements a number guessing game using binary search and logarithmic scaling.

The Game struct represents the main game state with the following fields:
- Peek: Current number being evaluated
- Boundary: Range limits for the search
- Mode: Current game state (Spaceless/Boundless/Bounded)
- Query: Handles the current question state
- Logarithm: Tracks search progress
- QueryText: The current question for the player

Core Functions:
- updateQuery(): Updates the game's current question based on game mode
- update(): Updates game state and progression
- stepUp(): Handles progression when number is higher than current guess
- stepDown(): Handles progression when number is lower than or equal to current guess
- Step(): Main game progression function that processes player responses

Game Modes:
- Spaceless: Initial state when the game starts
- Boundless: Exploring the number space to establish boundaries
- Bounded: Operating within known boundaries

The game uses a binary search algorithm with logarithmic scaling for efficient number guessing.
JSON tags are included for serialization compatibility.

Usage:
    g := game.CreateGame()
    g.Step() // Advances game state
    // Process g.QueryText for player input
*/
# Int Deduce

A number guessing game implementation in Go that uses binary search and logarithmic scaling to efficiently guess a number that a player is thinking of.

## How it Works

The game uses a smart algorithm that:
1. Narrows down the possible range using binary search
2. Uses logarithmic scaling for initial guesses
3. Asks the player if their number is more than the current guess
4. Adjusts the boundary based on the player's response

### Game States
The game uses three main states managed by the `iot` package:
- `Spaceless`: Initial state when the game starts
- `Boundless`: Exploring the number space to establish boundaries
- `Bounded`: Operating within known boundaries

### Core Components
- `Game`: Main structure containing game state and logic
    - `Peek`: Current number being evaluated
    - `Boundary`: Range limits for the search
    - `Stepper`: Current game state (Spaceless/Boundless/Bounded)
    - `Query`: Handles the current question state
    - `Logarithm`: Tracks search progress
    - `QueryText`: The current question for the player

### Features
- Efficient binary search implementation
- Dynamic boundary adjustment
- Logarithmic scaling for optimal guessing
- Pretty number formatting
- JSON-compatible structure

## Installation
```bash
go get github.com/renniemaharaj/int-deduce-go
```

## Usage
Import the packages:
```go
import (
        "github.com/renniemaharaj/int-deduce-go/pkg/game"
        "github.com/renniemaharaj/int-deduce-go/pkg/iot"
)
```

Create and use a game instance:
```go
func main() {
        g := game.CreateGame() // Creates game with default settings
        
        // Game loop example
        for {
                g.Step() // Advance game state
                fmt.Printf("Query: %s\n", g.QueryText)
                
                // Handle player input
                // '>' for numbers greater than current guess
                // '<' for numbers less than or equal to current guess
        }
}
```
