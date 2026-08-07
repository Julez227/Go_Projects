package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName[]ElementType

	// var numbers []int
	// var numbers1 []int = []int{1, 2, 3,}

	// numbers2 := []int{9, 8, 7,}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slicecopy:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)

	}

	fmt.Println("Element at index 3 of slice1", slice1[3])

	slice1[3] = 50
	fmt.Println("Element at index 3 of slice1 after modification", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		interLen := i + 1
		twoD[i] = make([]int, interLen)
		for j := 0; j < interLen; j++ {
			twoD[i][j] = i + j
			// fmt.Printf("twoD[%d][%d] = %d\n", i, j, twoD[i][j])
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice indiex of %d\n", i+j, i, j)
		}
	}

	fmt.Println("Two Dimensional Slice:", twoD)

	// slice[low:high] - slice from index low to high-1
	slice2 := slice1[2:4]
	fmt.Println("Slice2:", slice2)

	fmt.Println("The capacity of slice2 is:", cap(slice2))
	fmt.Println("The len of slice2 is:", len(slice2))

}
