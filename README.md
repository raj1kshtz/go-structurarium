# go-structurarium

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)](https://golang.org/dl/)
[![Tests](https://github.com/raj1kshtz/go-structurarium/workflows/Tests/badge.svg)](https://github.com/raj1kshtz/go-structurarium/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go library providing thread-safe, generic data structures with clean, easy-to-use APIs. Built with Go generics, this library offers type-safe implementations of common data structures including stacks, queues, vectors, collections, and hash maps.

## Features

- ✅ **Type-Safe Generics**: Leverages Go 1.18+ generics for compile-time type safety
- ✅ **Thread-Safe**: All data structures use channels for safe concurrent access
- ✅ **Well-Tested**: Comprehensive test coverage for all implementations
- ✅ **Clean API**: Simple, intuitive wrapper functions for all operations
- ✅ **Zero Dependencies**: Only uses standard library (except testing utilities)

## Data Structures

### Stack
A LIFO (Last-In-First-Out) data structure with the following operations:
- `Push(value)` - Add element to top
- `Pop()` - Remove and return top element
- `Peek()` - View top element without removing
- `Size()`, `IsEmpty()`, `Clear()`

### Queue
A FIFO (First-In-First-Out) data structure with operations:
- `Enqueue(value)` - Add element to rear
- `Dequeue()` - Remove and return front element
- `Peek()` - View front element
- `Size()`, `IsEmpty()`, `Clear()`, `ToArray()`

### Vector
A dynamic array with indexed access:
- `Add(value)`, `AddAt(index, value)` - Insert elements
- `Get(index)`, `Set(index, value)` - Access elements
- `RemoveAt(index)` - Remove element
- `Size()`, `IsEmpty()`, `Clear()`, `ToArray()`
- `EnsureCapacity()`, `TrimToSize()` - Capacity management

### Collection
A generic collection supporting various operations:
- `Add(value)`, `AddAll(values)` - Add elements
- `Remove(value)`, `RemoveAll(values)` - Remove elements
- `Contains(value)`, `ContainsAll(values)` - Check membership
- `RetainAll(values)` - Keep only specified elements
- `Size()`, `IsEmpty()`, `Clear()`, `ToArray()`

### HashMap
A key-value store with hash-based lookup:
- `Put(key, value)` - Insert or update
- `Get(key)` - Retrieve value
- `Remove(key)` - Delete entry
- `ContainsKey(key)` - Check key existence
- `Keys()`, `Values()` - Get all keys or values
- `Size()`, `IsEmpty()`, `Clear()`

## Installation

```bash
go get github.com/raj1kshtz/go-structurarium
```

## Quick Start

```go
import "github.com/raj1kshtz/go-structurarium/stack"

// Create a stack of integers
s := stack.NewWrapperStack[int]()
s.Push(10)
s.Push(20)
value, _ := s.Pop() // returns 20
```

For detailed usage examples, see [USAGE.md](USAGE.md)

## Requirements

- Go 1.18 or higher (for generics support)

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test ./... -v
```

## Project Structure

```
go-structurarium/
├── stack/          # Stack implementation
├── queue/          # Queue implementation
├── vector/         # Dynamic array implementation
├── collection/     # Generic collection implementation
├── maps/           # HashMap implementation
└── datastructure_helper/ # Example usage helpers
```

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting a Pull Request.

### Quick Start for Contributors

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes and add tests
4. Ensure all tests pass: `go test ./...`
5. Commit your changes: `git commit -m 'Add amazing feature'`
6. Push to your fork: `git push origin feature/amazing-feature`
7. Open a Pull Request

All PRs must:
- ✅ Pass all tests
- ✅ Include test coverage for new code
- ✅ Have at least one approval
- ✅ Have no unresolved conversations

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## Repository Setup

If you're the maintainer, see [.github/SETUP.md](.github/SETUP.md) for instructions on:
- Configuring branch protection
- Setting up automated testing and coverage reporting
- Configuring GitHub notifications

## License

This project is open source and available under the MIT License.

## Author

Created by [@raj1kshtz](https://github.com/raj1kshtz)