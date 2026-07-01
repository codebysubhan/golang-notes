package arrays

import "fmt"

func Learn(){
	var array [3]int

	array[0] = 0
	array[1] = 1
	array[2] = 2

	literal := [2]int{3,4}
	fmt.Println(array)
	fmt.Println(literal)


	// i := 0
	// for ; i < len(array) ; {
	// 	fmt.Println(array[i])
	// 	i++
	// }

	// is there any safer way of doing this iteration?

	for _ , val := range(array){
		// fmt.Println(_) // can't print it as we could do in python
		fmt.Println(val)
	}
}