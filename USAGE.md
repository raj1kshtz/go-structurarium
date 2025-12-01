# Usage Guide

This guide demonstrates how to use go-structurarium in your Go projects.

## Table of Contents

- [Installation](#installation)
- [Stack](#stack)
- [Queue](#queue)
- [Vector](#vector)
- [Collection](#collection)
- [HashMap](#hashmap)
- [Error Handling](#error-handling)
- [Custom Types](#custom-types)

## Installation

Add the library to your project:

```bash
go get github.com/raj1kshtz/go-structurarium
```

## Stack

A thread-safe, generic LIFO (Last-In-First-Out) stack implementation.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/stack"
)

func main() {
    // Create a new stack of integers
    s := stack.NewWrapperStack[int]()
    
    // Push elements
    s.Push(10)
    s.Push(20)
    s.Push(30)
    
    // Check size
    fmt.Println("Size:", s.Size()) // Output: Size: 3
    
    // Peek at top element
    top, _ := s.Peek()
    fmt.Println("Top:", top) // Output: Top: 30
    
    // Pop elements
    value, _ := s.Pop()
    fmt.Println("Popped:", value) // Output: Popped: 30
    
    // Check if empty
    fmt.Println("Is empty:", s.IsEmpty()) // Output: Is empty: false
    
    // Clear the stack
    s.Clear()
}
```

### String Stack Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/stack"
)

func main() {
    s := stack.NewWrapperStack[string]()
    
    s.Push("first")
    s.Push("second")
    s.Push("third")
    
    for !s.IsEmpty() {
        value, _ := s.Pop()
        fmt.Println(value)
    }
    // Output:
    // third
    // second
    // first
}
```

## Queue

A thread-safe, generic FIFO (First-In-First-Out) queue implementation.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/queue"
)

func main() {
    // Create a new queue
    q := queue.NewGenericQueueWrapper[int]()
    
    // Enqueue elements
    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)
    
    // Check size
    fmt.Println("Size:", q.Size()) // Output: Size: 3
    
    // Peek at front element
    front := q.Peek()
    fmt.Println("Front:", front) // Output: Front: 10
    
    // Dequeue elements
    value, success := q.Dequeue()
    if success {
        fmt.Println("Dequeued:", value) // Output: Dequeued: 10
    }
    
    // Convert to array
    arr := q.ToArray()
    fmt.Println("Array:", arr) // Output: Array: [20 30]
    
    // Clear the queue
    q.Clear()
}
```

### Processing Jobs Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/queue"
)

type Job struct {
    ID   int
    Name string
}

func main() {
    jobQueue := queue.NewGenericQueueWrapper[Job]()
    
    // Add jobs to queue
    jobQueue.Enqueue(Job{ID: 1, Name: "Process Payment"})
    jobQueue.Enqueue(Job{ID: 2, Name: "Send Email"})
    jobQueue.Enqueue(Job{ID: 3, Name: "Generate Report"})
    
    // Process jobs
    for !jobQueue.IsEmpty() {
        job, _ := jobQueue.Dequeue()
        fmt.Printf("Processing: %s (ID: %d)\n", job.Name, job.ID)
    }
}
```

## Vector

A thread-safe, generic dynamic array with indexed access.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/vector"
)

func main() {
    // Create a vector with initial capacity
    v := vector.NewWrapperVector[int](10)
    
    // Add elements
    v.Add(10)
    v.Add(20)
    v.Add(30)
    
    // Insert at specific index
    v.AddAt(1, 15) // Insert 15 at index 1
    
    // Get element at index
    value, _ := v.Get(1)
    fmt.Println("Value at index 1:", value) // Output: Value at index 1: 15
    
    // Update element
    v.Set(2, 25)
    
    // Remove element at index
    v.RemoveAt(0)
    
    // Get all elements as array
    arr := v.ToArray()
    fmt.Println("Vector:", arr) // Output: Vector: [15 25 30]
    
    // Size and capacity operations
    fmt.Println("Size:", v.Size())
    v.EnsureCapacity(20)
    v.TrimToSize()
}
```

### Iteration Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/vector"
)

func main() {
    v := vector.NewWrapperVector[string]()
    
    v.Add("apple")
    v.Add("banana")
    v.Add("cherry")
    v.Add("date")
    
    // Iterate through vector
    for i := 0; i < v.Size(); i++ {
        value, _ := v.Get(i)
        fmt.Printf("Index %d: %s\n", i, value)
    }
}
```

## Collection

A thread-safe, generic collection supporting set-like operations.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/collection"
)

func main() {
    // Create a collection
    c := collection.NewGenericCollectionWrapper[int]()
    
    // Add single element
    c.Add(10)
    c.Add(20)
    
    // Add multiple elements
    c.AddAll([]int{30, 40, 50})
    
    // Check if contains
    fmt.Println("Contains 30:", c.Contains(30)) // Output: Contains 30: true
    
    // Check if contains all
    fmt.Println("Contains all [20, 30]:", c.ContainsAll([]int{20, 30})) // Output: true
    
    // Remove element
    c.Remove(10)
    
    // Remove multiple elements
    c.RemoveAll([]int{40, 50})
    
    // Size
    fmt.Println("Size:", c.Size())
    
    // Convert to array
    arr := c.ToArray()
    fmt.Println("Collection:", arr)
}
```

### Set Operations Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/collection"
)

func main() {
    c := collection.NewGenericCollectionWrapper[string]()
    
    // Add elements
    c.AddAll([]string{"apple", "banana", "cherry", "date", "elderberry"})
    
    // Retain only specific elements (intersection)
    c.RetainAll([]string{"banana", "date", "fig"})
    
    // Result will be: ["banana", "date"]
    fmt.Println("After retain:", c.ToArray())
    
    // Clear collection
    c.Clear()
    fmt.Println("Is empty:", c.IsEmpty()) // Output: Is empty: true
}
```

## HashMap

A thread-safe, generic hash map implementation with key-value storage.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/maps"
)

func main() {
    // Create a hash map
    m := maps.NewGenericHashMapWrapper[string, int]()
    
    // Put key-value pairs
    m.Put("one", 1)
    m.Put("two", 2)
    m.Put("three", 3)
    
    // Get value by key
    value, exists := m.Get("two")
    if exists {
        fmt.Println("Value for 'two':", value) // Output: Value for 'two': 2
    }
    
    // Check if key exists
    fmt.Println("Contains 'one':", m.ContainsKey("one")) // Output: true
    
    // Update existing key
    m.Put("two", 20) // Updates value to 20
    
    // Get all keys
    keys := m.Keys()
    fmt.Println("Keys:", keys)
    
    // Get all values
    values := m.Values()
    fmt.Println("Values:", values)
    
    // Remove a key
    removed := m.Remove("three")
    fmt.Println("Removed:", removed) // Output: Removed: true
    
    // Size
    fmt.Println("Size:", m.Size())
    
    // Clear map
    m.Clear()
}
```

### Using HashMap with Custom Types

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/maps"
)

type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Map user ID to User struct
    userMap := maps.NewGenericHashMapWrapper[int, User]()
    
    // Add users
    userMap.Put(1, User{ID: 1, Name: "Alice", Age: 30})
    userMap.Put(2, User{ID: 2, Name: "Bob", Age: 25})
    userMap.Put(3, User{ID: 3, Name: "Charlie", Age: 35})
    
    // Get user by ID
    user, exists := userMap.Get(2)
    if exists {
        fmt.Printf("User: %s (Age: %d)\n", user.Name, user.Age)
    }
    
    // Iterate over all users
    for _, id := range userMap.Keys() {
        user, _ := userMap.Get(id)
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
    }
}
```

### Configuration Map Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/maps"
)

func main() {
    config := maps.NewGenericHashMapWrapper[string, string]()
    
    // Store configuration
    config.Put("host", "localhost")
    config.Put("port", "8080")
    config.Put("database", "myapp")
    config.Put("username", "admin")
    
    // Retrieve configuration
    host, _ := config.Get("host")
    port, _ := config.Get("port")
    
    fmt.Printf("Server running on %s:%s\n", host, port)
    
    // Check if configuration exists
    if config.ContainsKey("api_key") {
        apiKey, _ := config.Get("api_key")
        fmt.Println("API Key:", apiKey)
    } else {
        fmt.Println("API Key not configured")
    }
}
```

## Error Handling

Most operations return errors that should be checked:

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/stack"
    "github.com/raj1kshtz/go-structurarium/vector"
)

func main() {
    // Stack error handling
    s := stack.NewWrapperStack[int]()
    
    value, err := s.Pop()
    if err != nil {
        fmt.Println("Error:", err) // Error: stack underflow
    }
    
    // Vector error handling
    v := vector.NewWrapperVector[int]()
    
    _, err = v.Get(5)
    if err != nil {
        fmt.Println("Error:", err) // Error: index out of bounds
    }
    
    err = v.AddAt(-1, 10)
    if err != nil {
        fmt.Println("Error:", err) // Error: index out of bounds
    }
}
```

## Custom Types

All data structures support any type that satisfies their constraints:

- **Stack** and **Vector**: Work with any type (`any`)
- **Queue** and **Collection**: Work with comparable types
- **HashMap**: Keys must be comparable, values can be any type

### Example with Custom Struct

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/stack"
)

type Task struct {
    ID          int
    Description string
    Priority    int
}

func main() {
    taskStack := stack.NewWrapperStack[Task]()
    
    taskStack.Push(Task{ID: 1, Description: "Write code", Priority: 1})
    taskStack.Push(Task{ID: 2, Description: "Review PR", Priority: 2})
    taskStack.Push(Task{ID: 3, Description: "Fix bugs", Priority: 3})
    
    // Process tasks in LIFO order
    for !taskStack.IsEmpty() {
        task, _ := taskStack.Pop()
        fmt.Printf("Task %d: %s (Priority: %d)\n", 
                   task.ID, task.Description, task.Priority)
    }
}
```

## Performance Tips

1. **Preallocate capacity** when the size is known:
   ```go
   v := vector.NewWrapperVector[int](1000) // Preallocate for 1000 elements
   ```

2. **Use appropriate data structure**:
   - Use Stack for LIFO operations
   - Use Queue for FIFO operations
   - Use Vector for indexed access
   - Use Collection for set-like operations
   - Use HashMap for key-value lookups

3. **Batch operations** when possible:
   ```go
   c := collection.NewGenericCollectionWrapper[int]()
   c.AddAll([]int{1, 2, 3, 4, 5}) // Better than 5 individual Add() calls
   ```

## Thread Safety

All data structures are thread-safe by default, using channels for synchronization. You can safely use them across multiple goroutines:

```go
package main

import (
    "fmt"
    "sync"
    "github.com/raj1kshtz/go-structurarium/stack"
)

func main() {
    s := stack.NewWrapperStack[int]()
    var wg sync.WaitGroup
    
    // Multiple goroutines pushing to stack
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            s.Push(val)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("Stack size:", s.Size())
}
```

## Additional Examples

For more examples, see the `datastructure_helper` package in the repository, which contains helper functions demonstrating various use cases.

## Getting Help

If you encounter issues or have questions:
- Check the [README.md](README.md) for general information
- Review the test files for each data structure (e.g., `stack_test.go`, `queue_test.go`)
- Open an issue on GitHub

## Contributing

Contributions are welcome! Please see the [README.md](README.md) for contribution guidelines.

