package main

import (
	"fmt"
	"reflect"
)

func Hello(){ // the function with a capital letter can be use outside the package but if we don't want that then we use lowercase letters
	fmt.Println("Hello")
}

func Add(x int, y int) (bool, int) {
	x = 4
	// fmt.Println(x+y)
	// return true, x+y\
	if x+y != 0{
		return true, x+y
	} else {
		return false, x+y
	}
}

func main (){
	Hello()
	two := 2
	bl,out := Add(two, 3)
	fmt.Println(reflect.TypeOf(bl), reflect.TypeOf(out))
	fmt.Println(out, "<-- Returned Value", bl, "<-- boolean value")
	fmt.Println("two remains 2 because it is copied to x at the runtime so x and two are different varaibles here: ", two)
}