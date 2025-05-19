# Go Learning

This repository contains Go programming examples and exercises covering various concepts and implementations.

## Project Structure

```
.
├── basics/           # Basic Go programming concepts
├── concurrency/      # Concurrency examples
│   ├── mutex_example.go    # Mutex implementation
│   └── waitgroup.go       # WaitGroup examples
├── datastructures/   # Data structure implementations
│   ├── maps.go            # Map operations and examples
│   └── ...
├── examples/         # Example programs
│   ├── banking/           # Banking system examples
│   │   ├── demo.go       # Basic banking operations
│   │   └── teller.go     # Concurrent banking operations
│   ├── greeting/         # Greeting package examples
│   └── ...              # Other example programs
├── greeting/         # Greeting package implementation
├── mathpkg/         # Math operations package
├── patterns/        # Design pattern implementations
├── pointers/        # Pointer usage examples
└── main.go          # Main application entry point
```

## Key Components

### Data Structures
- Implementation of various data structures including maps
- Examples of Go's built-in data types and their operations

### Concurrency
- Mutex examples for safe concurrent access
- WaitGroup demonstrations for goroutine synchronization
- Channel usage and patterns

### Banking Examples
- Basic banking operations demonstration
- Thread-safe concurrent banking operations
- Account management and transaction handling

### Other Examples
- Greeting package usage and implementation
- Mathematical operations
- Pointer manipulation examples

## Getting Started

1. Clone the repository
2. Make sure you have Go 1.22 or later installed
3. Run examples:
   ```bash
   # Run main program
   go run main.go

   # Run specific examples
   go run examples/banking/demo.go
   go run examples/banking/teller.go
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
