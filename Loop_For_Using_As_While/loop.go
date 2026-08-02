package main

import "fmt"

func main() {

	// i := 1
	// for i <= 5 {
	// 	fmt.Println("interation:", i)
	// 	i++
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number:", num)
		num++ // ++ increment operator increases value by 1 and -- decrement operator, decreases value by 1
	}

}
