package datastructure_helper

import (
	"fmt"

	"github.com/raj1kshtz/go-structurarium/vector"
)

func VectorWrapperHelper() {
	intVector := vector.NewWrapperVector[int]()

	err := intVector.Add(10)
	if err != nil {
		fmt.Println("Error while adding to vector ", err)
	}

	err = intVector.AddAt(1, 20)
	if err != nil {
		fmt.Println("Error while adding to vector ", err)
	}

	err = intVector.RemoveAt(1)
	if err != nil {
		fmt.Println("Error while removing from vector ", err)
	}

	err = intVector.Set(1, 20)
	if err != nil {
		fmt.Println("Error while setting to vector ", err)
	}

	size := intVector.Size()
	fmt.Println("Vector size: ", size)

	isEmpty := intVector.IsEmpty()
	fmt.Println("Vector isEmpty: ", isEmpty)

	data := intVector.ToArray()
	fmt.Println("Vector toArray: ", data)

	err = intVector.TrimToSize()
	if err != nil {
		fmt.Println("Error while trimming to size ", err)
	}
}
