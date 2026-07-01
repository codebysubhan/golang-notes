package main

import (
	"fmt"
)

func Add(x *int, y *int) int {
	*x = 0
	return *x + *y
}

func main(){
	// & symbol for actually referencing the variable by its address
	value := 1
	fmt.Println(value, &value)

	// & *int --> type of the variable the pointer points to
	// *int pointer to an int type variable

	var ptr *int = &value
	fmt.Println(ptr, *ptr) // ptr here is the memory address of the value variable and *ptr is derefrencing the variable, dereferencing is just getting the value of that variable on that location


	val1 := 2
	val2 := 3
	fmt.Println(Add(&val1, &val2)) // we're passing the memory address because later we're dereferencing them
	fmt.Println("After the Add() function got called, val1 = ", val1)
}