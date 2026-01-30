# Usage Guide

This guide demonstrates how to use go-structurarium in your Go projects.

## Table of Contents

- [Installation](#installation)
- [Stack](#stack)
- [Queue](#queue)
- [Vector](#vector)
- [Collection](#collection)
- [HashMap](#hashmap)
- [Graph](#graph)
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


## Graph (Undirected & Directed)

A thread-safe, generic graph data structure supporting both undirected and directed graphs.

### Basic Usage (Undirected)

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/graph"
)

func main() {
    g := graph.NewUndirectedGraph[int, string]()
    g.AddVertex(1)
    g.AddVertex(2)
    g.AddEdge(1, 2, "edge-label")
    fmt.Println("Has edge 1-2:", g.HasEdge(1, 2)) // Output: true
    fmt.Println("Neighbors of 1:", g.Neighbors(1)) // Output: [2]
    g.RemoveEdge(1, 2)
    fmt.Println("Has edge 1-2 after removal:", g.HasEdge(1, 2)) // Output: false
}
```

### Basic Usage (Directed)

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/graph"
)

func main() {
    g := graph.NewDirectedGraph[int, string]()
    g.AddVertex(1)
    g.AddVertex(2)
    g.AddEdge(1, 2, "edge-label")
    fmt.Println("Has edge 1->2:", g.HasEdge(1, 2)) // Output: true
    fmt.Println("Has edge 2->1:", g.HasEdge(2, 1)) // Output: false
    fmt.Println("Neighbors of 1:", g.Neighbors(1)) // Output: [2]
    fmt.Println("Neighbors of 2:", g.Neighbors(2)) // Output: []
}
```

### Basic Usage (Wrapper)

```go
package main

import (
    "github.com/raj1kshtz/go-structurarium/graph"
)

func main() {
    gw := graph.NewGraphWrapper[int, string]()
    gw.AddVertex(1)
    gw.AddVertex(2)
    gw.AddEdge(1, 2, "edge-label")
    // ...
}
```

## Tree (N-ary Tree)

A generic tree structure where each node can have any number of children.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    // Create a tree with root value 1
    t := tree.NewTreeWrapper[int](1)
    
    // Add children to root
    t.Insert(1, 2)
    t.Insert(1, 3)
    t.Insert(1, 4)
    
    // Add grandchildren
    t.Insert(2, 5)
    t.Insert(2, 6)
    t.Insert(3, 7)
    
    fmt.Println("Tree size:", t.Size())        // Output: 7
    fmt.Println("Tree height:", t.Height())    // Output: 3
    
    // Check if value exists
    fmt.Println("Contains 5:", t.Search(5))    // Output: true
    fmt.Println("Contains 99:", t.Search(99))  // Output: false
}
```

### Tree Traversals

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    t := tree.NewTreeWrapper[int](1)
    t.Insert(1, 2)
    t.Insert(1, 3)
    t.Insert(2, 4)
    t.Insert(2, 5)
    t.Insert(3, 6)
    
    // Pre-order: root, then children
    fmt.Println("Pre-order:", t.PreOrder())
    // Output: [1 2 4 5 3 6]
    
    // Post-order: children, then root
    fmt.Println("Post-order:", t.PostOrder())
    // Output: [4 5 2 6 3 1]
    
    // Level-order: breadth-first
    fmt.Println("Level-order:", t.LevelOrder())
    // Output: [1 2 3 4 5 6]
}
```

### File System Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    // Create a file system structure
    fs := tree.NewTreeWrapper[string]("root")
    
    // Add directories
    fs.Insert("root", "home")
    fs.Insert("root", "usr")
    fs.Insert("root", "var")
    
    // Add subdirectories
    fs.Insert("home", "user1")
    fs.Insert("home", "user2")
    fs.Insert("usr", "bin")
    fs.Insert("usr", "lib")
    fs.Insert("var", "log")
    
    // Add files
    fs.Insert("user1", "documents")
    fs.Insert("user1", "downloads")
    
    // Display structure
    fmt.Println("File system structure (level-order):")
    for _, item := range fs.LevelOrder() {
        fmt.Println("-", item)
    }
    
    // Check if path exists
    if fs.Search("documents") {
        fmt.Println("documents found in the file system")
    }
    
    // Remove a directory and its contents
    fs.Remove("user2")
    fmt.Println("After removing user2, size:", fs.Size())
}
```

## Binary Search Tree (BST)

A self-organizing binary tree where left children are smaller and right children are larger than their parent.

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    // Create an empty BST
    bst := tree.NewBSTWrapper[int]()
    
    // Insert values
    bst.Insert(50)
    bst.Insert(30)
    bst.Insert(70)
    bst.Insert(20)
    bst.Insert(40)
    bst.Insert(60)
    bst.Insert(80)
    
    // Search for values
    fmt.Println("Contains 40:", bst.Search(40))   // Output: true
    fmt.Println("Contains 100:", bst.Search(100)) // Output: false
    
    // Get min and max
    min, _ := bst.Min()
    max, _ := bst.Max()
    fmt.Println("Min:", min) // Output: 20
    fmt.Println("Max:", max) // Output: 80
    
    // Get sorted values (in-order traversal)
    fmt.Println("Sorted:", bst.InOrder())
    // Output: [20 30 40 50 60 70 80]
}
```

### BST Operations

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    bst := tree.NewBSTWrapper[int]()
    
    // Insert multiple values
    values := []int{50, 30, 70, 20, 40, 60, 80}
    for _, v := range values {
        bst.Insert(v)
    }
    
    fmt.Println("Size:", bst.Size())     // Output: 7
    fmt.Println("Height:", bst.Height()) // Output: 3
    
    // Delete operations
    bst.Delete(20) // Delete leaf node
    bst.Delete(30) // Delete node with one child
    bst.Delete(50) // Delete node with two children
    
    fmt.Println("After deletions:", bst.InOrder())
    // Output: [40 60 70 80]
    
    // Validate BST property
    fmt.Println("Is valid BST:", bst.Validate()) // Output: true
}
```

### Sorted Data Management

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    // BST for maintaining sorted strings
    wordTree := tree.NewBSTWrapper[string]()
    
    // Insert words
    words := []string{"dog", "cat", "elephant", "ant", "bear"}
    for _, word := range words {
        wordTree.Insert(word)
    }
    
    // Get words in sorted order
    fmt.Println("Sorted words:", wordTree.InOrder())
    // Output: [ant bear cat dog elephant]
    
    // Find alphabetically first and last
    first, _ := wordTree.Min()
    last, _ := wordTree.Max()
    fmt.Println("First:", first) // Output: ant
    fmt.Println("Last:", last)   // Output: elephant
}
```

### Student Grades Example

```go
package main

import (
    "fmt"
    "github.com/raj1kshtz/go-structurarium/tree"
)

func main() {
    // BST for student scores
    scores := tree.NewBSTWrapper[int]()
    
    // Add test scores
    testScores := []int{85, 92, 78, 95, 88, 73, 90}
    for _, score := range testScores {
        scores.Insert(score)
    }
    
    // Get all scores in sorted order
    sorted := scores.InOrder()
    fmt.Println("Scores (sorted):", sorted)
    
    // Find median score (middle value)
    median := sorted[len(sorted)/2]
    fmt.Println("Median score:", median)
    
    // Get highest and lowest scores
    lowest, _ := scores.Min()
    highest, _ := scores.Max()
    fmt.Println("Range:", lowest, "-", highest)
    
    // Check if passing grade (60) exists
    if scores.Search(60) {
        fmt.Println("Someone got exactly 60")
    } else {
        fmt.Println("No one got exactly 60")
    }
}
```

### BST vs Tree Comparison

**When to use N-ary Tree:**
- Hierarchical data (file systems, org charts)
- No ordering required
- Variable number of children per node
- Tree traversal is more important than search

**When to use BST:**
- Need efficient search (O(log n) average)
- Need sorted data
- Range queries
- Finding min/max values frequently
- Exactly two children per node

## Additional Examples

For more examples, see the `datastructure_helper` package in the repository, which contains helper functions demonstrating various use cases.

## Getting Help

If you encounter issues or have questions:
- Check the [README.md](README.md) for general information
- Review the test files for each data structure (e.g., `stack_test.go`, `queue_test.go`)
- Open an issue on GitHub

## Contributing

Contributions are welcome! Please see the [README.md](README.md) for contribution guidelines.

