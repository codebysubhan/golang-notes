package structs

import "fmt"

// single-use variable with a temporary shape
var myStruct struct{ // struct is a built in type that could contain other types
	value int
	num float32
	str string
}

// creating a named struct, which acts as a reusable blueprint
// defining custom types using type keyword
type person struct { // to access this struct outside of this package we have to make the p of person capitalized as Person
	age int
	name string
}

func showUser(p person){
	fmt.Println(p)
}

func modifyVal(p *person){
	p.age = 35
	p.name = "charlie"
}

func LearnStruct(){
	// myStruct.value = 2
	// myStruct.num = 3.14
	// myStruct.str = "pi"
	// fmt.Println(myStruct)

	// fmt.Println(person) // will give an error

	var user person = person{30,"bob"}
	showUser(user)
	modifyVal(&user)
	showUser(user)

}