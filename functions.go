package main

import "fmt"

func functions() {
	x, y := 5, 6

	fmt.Println(add(x,y))

	num := 5
	fmt.Println(factorial(num))
}

// making a function that takes in two integers and returns an integer
func add(num1, num2 int) int {
	return num1+num2
}

// recursive function (where you call a function within itself)
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num -1)
	}
}
