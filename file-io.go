package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func file() {
	// creates a new file
	file, err := os.Create("sample.txt")

	// catches for errors
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This file was created using Go")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// displaying file string in terminal
	s1 := string(stream)
	fmt.Println(s1)
}
