package main

import "fmt"

func arrays() {
	// ways to initialise arrays
	// 1st way
	var EvenNum[3] int

	EvenNum[0] = 2
	EvenNum[1] = 4
	EvenNum[2] = 6
	fmt.Println(EvenNum[2])

	// 2nd way
	OddNum := [3]int{1, 3, 5}

	// looping through an array
	for _, value := range OddNum { // underscore means we're not using an iterator & we're looping in the range of the array
		fmt.Println(value)
	}

	// slicing arrays
	sliced := EvenNum[0:1]
	fmt.Println(sliced)

	numSlice := []int{5, 4, 3, 2, 1}

	// empty slice with length of 5 and capacity of 10
	slice2 := make([]int, 5, 10)

	// copy one slice to another
	copy(slice2, numSlice)
	fmt.Println(slice2)

	// appending values to a slice
	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println(slice3)
}
