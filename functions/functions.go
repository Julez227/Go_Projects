package main

import "fmt"

func main() {

	// // func <name> (<parameters>) <return type> {
	// // code block
	// // return value
	// // }

	// // sum := add(1, 2)
	// fmt.Println(add(2, 3))

	// greet := func() {
	// 	fmt.Println("Hello from anonymous function greet")
	// }

	// greet()

	// func() {
	// 	fmt.Println("Hello from anonymous function")
	// }() // calling the anonymous function

	// operation := add

	// restul := operation(3, 5)
	// fmt.Println(restul)

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5+3 =", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6*2 =", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

// function that takes a function as a argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}

}
