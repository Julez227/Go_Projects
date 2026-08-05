package main

import "fmt"

func main() {

	// Switch statement in go is (switch case default) (fallthrough allows the code to move to next case )
	// switch expression {
	// case value 1:
	// 	// code to be executed if expression equals value1
	// case value 2:
	// 	// code to be executed if expression equals value2
	// case value 3:
	// 	// code to be executed if expression equals value3
	// default:
	// code to be executed if expression does not match any value
	// }

	// fruit := "Cherry"

	// switch fruit {
	// case "Apple":
	// 	fmt.Println("Its an apple.")
	// case "banna":
	// 	fmt.Println("its a banna.")
	// default:
	// 	fmt.Println("Unknown Fruit")
	// }

	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("Its a weekday.")
	// case "Sunday":
	// 	fmt.Println("Its a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("number is 2")
	// default:
	// 	fmt.Println("Not Two")
	// }
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

// type switch
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Its an interger")
	case float64:
		fmt.Println("Its float")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Unknown Type")
	}
}
