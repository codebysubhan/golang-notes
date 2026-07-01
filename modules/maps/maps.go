package maps

import (
	"fmt"
	"reflect"
)

func LearnMaps(){
	// myMap := make(map[int]string) // int is the key and string is the value
	// myMap[5] = "five"
	// myMap[1] = "one"
	// fmt.Println(myMap)

	// we can do it using the map literal as well
	myMap := map[string]int{"one":1, "five":5}
	fmt.Println(myMap, reflect.TypeOf(myMap))


	// accessing something that doesnt exists returns 0 but what if we need to check if zero is actually there against a key?
	fmt.Println(myMap["asb"])

	// value exists
	ok, val := myMap["one"]
	fmt.Println(ok, val)

	// value doesn't exists
	ok1, val1 := myMap["aksjdh"]
	fmt.Println(ok1, val1)

	// how to delete a key, value pair?
	// use delete
	// delete(myMap, "one")
}