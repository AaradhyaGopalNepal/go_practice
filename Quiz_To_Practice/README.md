# Go Programming Challenge: 2-Month Learning Journey 🚀

## Overview
This repository contains a comprehensive 8-week (2-month) Go programming challenge designed to take you from basics to advanced concepts through hands-on, test-driven development. Each week focuses on specific Go concepts with challenging exercises that build upon previous knowledge.

## Structure
```
go-quiz/
├── week1_basics/           # Variables, Types, Control Flow
├── week2_functions/        # Functions, Defer, Panic, Recover
├── week3_collections/      # Arrays, Slices, Maps
├── week4_structs/          # Structs and Methods
├── week5_interfaces/       # Interfaces and Polymorphism
├── week6_concurrency/      # Goroutines and Channels
├── week7_advanced_concurrency/  # Advanced Sync Patterns
└── week8_projects/         # Final Project: CLI Task Runner
```

## How to Use This Repository

### Initial Setup
```bash
# Clone the repository
git clone <your-repo-url>
cd go-quiz

# Install dependencies
go mod download

# Verify setup
go test ./... -v
```

### Working Through Challenges

1. **Start with Week 1** and progress sequentially
2. Each challenge file has functions/methods **commented out**
3. **Uncomment** the code block you want to implement
4. **Write your implementation**
5. **Run the tests** to verify your solution
6. Tests have a `t.Skip()` line - **remove it** when ready to test

### Example Workflow
```bash
# Navigate to a week's folder
cd week1_basics

# Open a challenge file
# Uncomment the function you want to implement
# Write your implementation

# Run specific test
go test -v -run TestConvertAndProcess

# Run all tests for the file
go test -v 01_type_converter_test.go 01_type_converter.go

# Run all tests for the week
go test -v ./...
```

## Weekly Breakdown

### Week 1: Variables, Types, and Control Flow
**Files:** `01_type_converter.go`, `02_control_flow_master.go`, `03_operators_challenge.go`

**Concepts Covered:**
- Type conversions and assertions
- Complex control flow patterns
- Bitwise operations
- Logical operators
- Pattern matching

**Key Challenges:**
- Build a type analyzer that detects and converts between types
- Implement FizzBuzzPlus with multiple rules
- Create a pattern matcher for strings
- Master bitwise manipulations

### Week 2: Functions and Error Handling
**Files:** `01_function_fundamentals.go`, `02_defer_panic_recover.go`

**Concepts Covered:**
- Multiple return values
- Variadic functions
- Closures and higher-order functions
- Defer statements
- Panic and recover mechanisms

**Key Challenges:**
- Build an advanced calculator with multiple returns
- Create a counter generator using closures
- Implement resource management with defer
- Handle panics gracefully with recover

### Week 3: Arrays, Slices, and Maps
**Files:** `01_slice_mastery.go`, `02_map_expertise.go`

**Concepts Covered:**
- Slice operations and capacity management
- 2D slices and matrix operations
- Map operations and algorithms
- Graph representations
- LRU cache implementation

**Key Challenges:**
- Implement slice rotation and manipulation algorithms
- Build a sliding window algorithm
- Create frequency analyzers
- Implement graph traversal algorithms
- Build an LRU cache from scratch

### Week 4: Structs and Methods
**Files:** `01_struct_methods.go`

**Concepts Covered:**
- Struct composition and embedding
- Method receivers
- Custom data structures
- Object-oriented design in Go

**Key Challenges:**
- Build a banking system with accounts and transactions
- Create a library management system
- Implement shape hierarchy with embedded structs
- Build an ordered map data structure
- Create an employee hierarchy with composition

### Week 5: Interfaces and Polymorphism
**Files:** `01_interface_design.go`

**Concepts Covered:**
- Interface design
- Multiple interface implementations
- Factory patterns
- Strategy pattern
- Dependency injection

**Key Challenges:**
- Build multiple storage backends (Memory, File, Cache)
- Create a database abstraction layer
- Implement a multi-format serializer
- Build a payment processor with multiple providers
- Create a flexible logging system

### Week 6: Concurrency (Goroutines and Channels)
**Files:** `01_goroutines_channels.go`

**Concepts Covered:**
- Goroutines
- Channels (buffered and unbuffered)
- Select statements
- Worker pools
- Fan-in/Fan-out patterns

**Key Challenges:**
- Build a parallel processor with worker pools
- Implement pipeline processing
- Create a rate limiter
- Build producer-consumer pattern
- Implement a broadcaster system

### Week 7: Advanced Concurrency Patterns
**Files:** `01_sync_patterns.go`

**Concepts Covered:**
- Custom synchronization primitives
- Context package
- Advanced patterns (Circuit Breaker, Retry)
- Object pooling
- Futures and promises

**Key Challenges:**
- Implement a semaphore from scratch
- Build a custom read-write lock
- Create a barrier synchronization primitive
- Implement circuit breaker pattern
- Build a retry mechanism with exponential backoff

### Week 8: Final Project - CLI Task Runner
**Files:** `01_task_runner_cli.go`

**The Ultimate Challenge:** Build a complete distributed task runner CLI application

**Features to Implement:**
- Task scheduling with priority queue
- Concurrent execution with worker pools
- Dependency resolution
- Persistent storage
- Rate limiting and circuit breaking
- REST API server
- CLI interface with commands
- Plugin system
- Metrics and monitoring
- Real-time event streaming

This project integrates ALL concepts learned in previous weeks!

## Testing Guidelines

### Running Tests
```bash
# Run all tests (will show skipped)
go test ./... -v

# Run tests for specific week
go test ./week1_basics/... -v

# Run with coverage
go test ./... -cover

# Run specific test function
go test -v -run TestFunctionName ./week1_basics/

# Run tests with race detection
go test -race ./...
```

### Test Structure
Each test file includes:
- Comprehensive test cases
- Edge cases
- Performance considerations
- Clear assertions
- Helper functions when needed

## Tips for Success

### 1. Follow TDD (Test-Driven Development)
- Read the test first to understand requirements
- Write minimal code to make tests pass
- Refactor for better design

### 2. Progressive Learning
- Don't skip weeks - concepts build on each other
- Complete at least 80% of each week before moving on
- Review previous implementations when stuck

### 3. Best Practices
- Write clean, idiomatic Go code
- Use meaningful variable names
- Add comments for complex logic
- Handle errors properly
- Think about edge cases

### 4. When Stuck
- Review the test cases for hints
- Check Go documentation
- Use `go doc` command for standard library help
- Remember: Tests define the expected behavior

## Completion Checklist

### Week 1 ✅
- [ ] Type Converter
- [ ] Control Flow Master
- [ ] Operators Challenge

### Week 2 ✅
- [ ] Function Fundamentals
- [ ] Defer, Panic, Recover

### Week 3 ✅
- [ ] Slice Mastery
- [ ] Map Expertise

### Week 4 ✅
- [ ] Structs and Methods

### Week 5 ✅
- [ ] Interface Design

### Week 6 ✅
- [ ] Goroutines and Channels

### Week 7 ✅
- [ ] Sync Patterns

### Week 8 ✅
- [ ] Task Runner CLI (Final Project)

## Resources

### Official Documentation
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Blog](https://blog.golang.org/)

### Books
- "The Go Programming Language" by Donovan & Kernighan
- "Go in Action" by Kennedy, Ketelsen & St. Martin
- "Concurrency in Go" by Katherine Cox-Buday

### Online Resources
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://play.golang.org/)
- [Awesome Go](https://github.com/avelino/awesome-go)

## Troubleshooting

### Common Issues

**Import errors:**
```bash
go mod tidy
```

**Test timeouts:**
Some concurrency tests may need longer timeouts:
```bash
go test -timeout 30s ./...
```

**Race conditions:**
Always test concurrent code with race detector:
```bash
go test -race ./week6_concurrency/...
```

## Contributing
Feel free to:
- Add more test cases
- Suggest additional challenges
- Improve implementations
- Fix bugs

## License
This project is for educational purposes. Feel free to use and modify as needed for your learning journey.

---

## Final Notes

**Congratulations on starting this journey!** 🎉

This challenge set is designed to be difficult but rewarding. Each week builds essential Go programming skills that are used in real-world applications.

The Week 8 final project is intentionally comprehensive and time-consuming. It's designed to bring together everything you've learned into a production-grade application. Take your time with it - building it properly will solidify your Go expertise.

**Remember:**
- There's no shame in looking up documentation
- The best way to learn is by doing
- Errors are learning opportunities
- Persistence pays off

Good luck, and happy coding! 🚀