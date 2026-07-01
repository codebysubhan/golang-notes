package main

import (
	"fmt"
)

func main(){ // why I'm getting a warning saying that main is redeclared in this block and it's pointing towards the hello.go file however, the program runs.
	// conditions
	// if 1 == 1 {
	// 	fmt.Println("1 equals 1")
	// } else if 1 != 2 && "Hello" == "Hello" { // for this block to run the first one should fail
	// 	fmt.Println("Double equals")
	// }

	// loops
	for i:=0; i < 10; i++ {
		fmt.Println(i)
	}

	// we can use continue and break the similar way we use in python inside a loop
}