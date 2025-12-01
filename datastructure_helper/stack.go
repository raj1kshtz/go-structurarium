package datastructure_helper

import (
	"fmt"

	"github.com/raj1kshtz/go-structurarium/stack"
)

func StackWrapperHelper() {
	intStack := stack.NewWrapperStack[int]()
	err := intStack.Push(10)
	if err != nil {
		fmt.Println("Error while pushing to intStack:", err)
	}
	err = intStack.Push(20)
	if err != nil {
		fmt.Println("Error while pushing to intStack:", err)
	}

	value, err := intStack.Pop()
	if err == nil {
		fmt.Println("Popped element is :", value)
	} else {
		fmt.Println("Error popping from intstack:", err)
	}

	top, err := intStack.Peek()
	if err == nil {
		fmt.Println("Top element is :", top)
	}
	err = intStack.Clear()
	if err != nil {
		fmt.Println("Error while clearing intStack:", err)
	}
	fmt.Println(intStack.IsEmpty())

	personStack := stack.NewWrapperStack[Person]()

	person1 := Person{ID: 1, Name: "John"}
	person2 := Person{ID: 2, Name: "Doe"}

	err = personStack.Push(person1)
	if err != nil {
		fmt.Println("Error pushing John to stack:", err)
	}
	err = personStack.Push(person2)
	if err != nil {
		fmt.Println("Error pushing Doe to stack:", err)
	}

	poppedPerson, err := personStack.Pop()
	if err == nil {
		fmt.Printf("Popped Person: %+v\n", poppedPerson)
	} else {
		fmt.Println("Error popping from person stack:", err)
	}

	topPerson, err := personStack.Peek()
	if err == nil {
		fmt.Printf("Top Person: %+v\n", topPerson)
	} else {
		fmt.Println("Error getting top person:", err)
	}
	err = personStack.Clear()
	if err != nil {
		fmt.Println("Error while clearing personStack:", err)
	}
}
