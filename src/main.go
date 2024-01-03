package main

import "fmt"

func main() {
	// Area of a scuare
	const baseScuare = 10
	areaScuare := baseScuare * baseScuare
	fmt.Println("Area scuare: ", areaScuare)

	x := 10
	y := 50

	// addition
	result := x + y
	fmt.Println("Addition: ", result)

	// subtraction
	result = y - x
	fmt.Println("Subtraction: ", result)

	//Multiplication
	result = y * x
	fmt.Println("Multiplication: ", result)

	// division
	result = y / x
	fmt.Println("Division: ", result)

	// Module
	result = y % x
	fmt.Println("Module: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x++
	fmt.Println("Decremental: ", x)

}
