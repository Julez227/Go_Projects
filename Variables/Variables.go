package main

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var nam1= "Jane"

	// count := 10
	// lastName := "Smith"
	middleName := "Mayor"
	fmt.Println(middleName)
	// Default values
	// Numeric Types: 0
	// Boolean Types: false
	// String Types: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ----SCOPE

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}

// // ---- VARIABLES ----

// // "var" explicitly declares a variable with a type
// var age int = 30

// // Go can infer the type automatically if you skip it
// var city = "Cody"       // Go infers this is a string

// // ":=" is shorthand — declares AND assigns in one step
// // Only works INSIDE functions, not at package level
// count := 5

// // Variables declared with var/:= CAN be reassigned later
// age = 31          // valid — age was declared as a var, can change
// count = 10        // valid — same reason

// // Multiple variables can be declared together in a block
// var (
//     firstName string = "Julius"
//     lastName  string = "Williams"
//     isRemote  bool   = true
// )

// // If you declare a variable but don't assign a value,
// // Go automatically fills it with a "zero value" based on type:
// var score int      // defaults to 0
// var label string   // defaults to ""
// var active bool    // defaults to false
