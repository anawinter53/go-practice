package main

import "fmt"

func variables() {
	// when declaring type write it after variable name
	const pi float64 = 3.1415926535

	// walrus operator to declare w/ short-hand
	x, y := 14, 15

	fmt.Println(x, y)

	// prints out float in formatted print to 3dp
	fmt.Printf("%.3f \n", pi) 
}

