package main

import "fmt"

func control() {
	// using defer to delay execution of fxn until surrounding fxn returns
	defer FirstRun()
	SecondRun()

	// recover and panic
	fmt.Println(div(3,0))
	fmt.Println(div(3,5))
	demPanic()

	// passing unlimited arg to a fxn
	fmt.Println(addemup(10, 20, 30, 40, 50))
}

func FirstRun() { fmt.Println("I now execute second") }
func SecondRun() { fmt.Println("I now execute first") }

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1/num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}

func addemup(args ... int) int {
	sum := 0

	for _, value := range args {
		sum += value
	}
	return sum
}
