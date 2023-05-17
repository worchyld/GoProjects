package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// Simple addition
	// Uses shorthand for variables
	x := 5
	y := 7
	sum := x + y

	fmt.Println(sum)

	if sum > 10 {
		fmt.Println("More than 10")
	}
}