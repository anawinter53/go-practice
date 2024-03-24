package main

import "fmt"

func loops() {
	// for loop
	for i := 1; i <= 10; i ++ {
		fmt.Println(i)
	}
	
	// making a for loop with while loop style syntax
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// making an if statement for age
	age := 18

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You're too young to vote")
	}

	// now with a switch statement
	switch age {
		case 16: fmt.Println("You can vote in two years")
		case 18: fmt.Println("Great, you can vote now")
		default: fmt.Println("Apparently you don't have an age")
	}
}
