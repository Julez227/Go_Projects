package main

import (
	"fmt"
)

func main() {

	// /... Ellipsis

	// func funtionName(parameter1 typer1, param2, param3 ...type3 ) returnType{
	// function body
	// }

	// fmt.Println("sum of 1, 2, 3:", sum(1, 2, 3))
	statement, total := sum("The sum of 1,2,3 is", 1, 2, 3)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence3, total3 := sum("Sequence: ", numbers...)
	fmt.Println("Sequence: ", sequence3, "Total", total3)

}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
