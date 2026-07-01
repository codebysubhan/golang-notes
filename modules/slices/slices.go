package slices

import "fmt"

func Slices()  {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 15)
	fmt.Println(slice)
	slice = append(slice, 155)
	fmt.Println(slice)

	// slicing with size
	sliceWithSize := make([]int, 7) // can change int to string and it will print the default value of that and it is space
	fmt.Println(sliceWithSize) // will print all 0s

	// AI overview of array vs slice
	// arr := [5]int{1, 2, 3, 4, 5}  // Array (Fixed size of 5)
	// arr := [...]int{1, 2, 3}      // Array (Go counts the elements and fixes it at 3)
	// slc := []int{1, 2, 3}         // Slice (Dynamic size, no number specified)

	for _, val := range(slice){
		fmt.Println(val)
	}
}
