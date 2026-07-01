// main is the entry point package and its suposed to be run directly
package main

// import only the packages that we need
import (
	"fmt"
	"reflect"
)

// a go program always looks for the function named main
func main(){
	fmt.Println("Hello, Go!!!") // string
	fmt.Println('A') // runes are kept as numerical unicode
	
	var boolean bool = true // typed
	// boolean2 := true // untyped
	fmt.Println(boolean)

	// integers, floats
	// for addition + , subtraction - , multiplication * , dvision /
	// < > == != >= <=

	// how to check the type of the variable?
	fmt.Println(reflect.TypeOf(3.14))

	// variable declaration using typed way
	// var quantity int = 4
	// var min, max float64
	// var name string

	// min, max = 2, 10
	// name = "Buster"
	// compiler will complain if we declare/initialize something and doesn't use it

	// normally we use shorthand untyped way
	quantity := 4
	min, max := 3, 5
	fmt.Println(quantity, min, max)



}
// in go a string is doubled quoted and rune is single quoted
// rune is similar to the char as in python

