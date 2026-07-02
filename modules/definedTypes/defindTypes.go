package definedtypes

import "fmt"

type person struct{
	age int
	name string
}
// & generates a memory address (a pointer)
// * to get the actual value out of it
// custom method for a custom type the difference is that we included a receriver parameter with pointer (p *person)
func (p *person) changeUser(){ // without this pointer the actual value won't change
	p.age = 20
	p.name = "bobby"
}

func LearnDefinedType(){
	var user person = person{30, "bob"}
	user.changeUser()
	fmt.Println(user)
}