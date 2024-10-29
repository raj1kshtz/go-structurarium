package datastructure_helper

import (
	"fmt"
	"go-structurarium/queue"
)

func QueueWrapperHelper() {
	intQueue := queue.NewGenericQueue[int]()
	err := intQueue.Offer(20)
	if err != nil {
		fmt.Println("Error with offering to intQueue", err)
	}

	err = intQueue.AddFirst(10)
	if err != nil {
		fmt.Println("Error with adding First to intQueue", err)
	}

	err = intQueue.AddLast(30)
	if err != nil {
		fmt.Println("Error with adding Last to intQueue", err)
	}

	val, err := intQueue.Peek()
	if err != nil {
		fmt.Println("Error with peeking from intQueue", err)
	} else {
		fmt.Println("Value from intQueue", val)
	}

	val, err = intQueue.Poll()
	if err != nil {
		fmt.Println("Error with polling from intQueue", err)
	} else {
		fmt.Println("Value from intQueue", val)
	}

	val, err = intQueue.PollFirst()
	if err != nil {
		fmt.Println("Error with polling first from intQueue", err)
	} else {
		fmt.Println("Value from intQueue", val)
	}

	intQueue.Display()
	fmt.Println("intQueue Size", intQueue.Size())

	val, err = intQueue.PollLast()
	if err != nil {
		fmt.Println("Error with polling last from intQueue", err)
	} else {
		fmt.Println("Value from intQueue", val)
	}

	err = intQueue.Clear()
	if err != nil {
		fmt.Println("Error with clearing from intQueue", err)
	}

	fmt.Println("intQueue IsEmpty", intQueue.IsEmpty())

	personQueue := queue.NewGenericQueue[Person]()
	err = personQueue.Offer(Person{ID: 1, Name: "Alice"})
	if err != nil {
		fmt.Println("Error with offering to personQueue:", err)
	}

	err = personQueue.AddFirst(Person{ID: 2, Name: "Bob"})
	if err != nil {
		fmt.Println("Error with adding First to personQueue:", err)
	}

	err = personQueue.AddLast(Person{ID: 3, Name: "Charlie"})
	if err != nil {
		fmt.Println("Error with adding Last to personQueue:", err)
	}

	peekedPerson, err := personQueue.Peek()
	if err != nil {
		fmt.Println("Error with peeking from personQueue:", err)
	} else {
		fmt.Printf("Peeked person from personQueue: %+v\n", peekedPerson)
	}

	polledPerson, err := personQueue.Poll()
	if err != nil {
		fmt.Println("Error with polling from personQueue:", err)
	} else {
		fmt.Printf("Polled person from personQueue: %+v\n", polledPerson)
	}

	firstPerson, err := personQueue.PollFirst()
	if err != nil {
		fmt.Println("Error with polling first from personQueue:", err)
	} else {
		fmt.Printf("First person from personQueue: %+v\n", firstPerson)
	}

	personQueue.Display()
	fmt.Println("personQueue Size:", personQueue.Size())

	lastPerson, err := personQueue.PollLast()
	if err != nil {
		fmt.Println("Error with polling last from personQueue:", err)
	} else {
		fmt.Printf("Last person from personQueue: %+v\n", lastPerson)
	}

	err = personQueue.Clear()
	if err != nil {
		fmt.Println("Error with clearing personQueue:", err)
	}

	fmt.Println("personQueue IsEmpty:", personQueue.IsEmpty())
}
