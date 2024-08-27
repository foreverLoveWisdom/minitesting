package main

import (
	"fmt"
)

type T struct {
	failed bool
	output func(string, ...interface{}) (int, error) // Match fmt.Printf's signature
}

func main() {
	// Create an instance of T and assign fmt.Printf to output
	defaultT := &T{
		failed: false,
		output: fmt.Printf, // This is now valid
	}

	// Using the output function
	defaultT.output("The result is %d\n", 42)
}
