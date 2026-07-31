package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// used for naming Structs, interfaces, enums. distinguishes type names from variables and functions

	// snake_case
	// Eg. user_id, first_name, http_request
	// used for naming variables, constants, and file names.

	// UPPERCASE
	// ALL CAPS
	// Used exclusively for naming constants
	// ensures that constants stand out and thier immutability is emphasized

	// mixedCase
	// Eg. javaScript, htmlDocument, isValid
	// used to name variables or certain identifiers

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
