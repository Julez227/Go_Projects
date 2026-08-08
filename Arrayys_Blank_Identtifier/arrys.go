package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banna", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	originalArry := [3]int{1, 2, 3}
	copiedArry := originalArry

	copiedArry[0] = 100

	fmt.Println("Original arry:", originalArry)
	fmt.Println("Copied arry:", copiedArry)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	//_ underscore is blank idenifier, used to store unused
}
	