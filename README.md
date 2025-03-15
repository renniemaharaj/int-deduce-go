# Int Deduce

A number guessing game implementation in Go that uses binary search and logarithmic scaling to efficiently guess a number that a player is thinking of.

## How it Works

The game uses a smart algorithm that:
1. Narrows down the possible range using binary search
2. Uses logarithmic scaling for initial guesses
3. Asks the player if their number is more than the current guess
4. Adjusts the boundary based on the player's response

### Game States
- `None`: Initial state
- `Is`: Asking if the number is more than the current guess
- `Not`: Processing negative responses

### Features
- Efficient number searching
- JSON-compatible structure
- Pretty number printing
- Dynamic query text updates
- Boundary management

## Installation
```bash
go get github.com/renniemaharaj/int-deduce-go
```

## Usage
Import the required packages:
```go
import (
    "github.com/renniemaharaj/int-deduce-go/pkg/game"
    "github.com/renniemaharaj/int-deduce-go/pkg/iotaTypes"
)
```

Create and run a game:
```go
func main() {
    g := game.CreateGame()

    for {
        g.Move()
        fmt.Printf("\n--- Game Status ---\n")
        fmt.Printf("Query: %s\n", g.QueryText)
        fmt.Printf("Boundary: [%d, %d]\n", g.Boundary.Start, g.Boundary.End)
        fmt.Printf("Steps Remaining: %.2f\n", g.Logarithm)
        fmt.Printf("State: %v\n", g.State)
        fmt.Println("------------------")
        fmt.Println("Enter '>' (More), '<' (Less), or '/' (Exit):")

        var response string
        if _, err := fmt.Scanln(&response); err != nil {
            fmt.Println("Error:", err)
            continue
        }

        switch response {
        case ">":
            g.Query.SetConfirmed(iotaTypes.Is)
        case "<":
            g.Query.SetConfirmed(iotaTypes.Not)
        case "/":
            return
        default:
            fmt.Println("Invalid input. Use '>', '<', or '/'")
        }
    }
}
```

The game will automatically adjust its guesses based on player responses until it finds the correct number.
