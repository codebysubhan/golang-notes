package encapsulation

import (
	// "log"
	"fmt"
	"example.com/modules/adult"
)


type child struct{
	age int
	name string
	adult.Adult // this is called embedding and it will let the child type use getters and setters of the adult type
}


func Encap(){
	// parent := adult.Adult{30 , "bob"}
	// err := parent.SetAge(32)
	// // fmt.Println(err, "<--- printed error at stdout")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(parent.age)
	// out := parent.GetAge()
	// fmt.Println(out)
	kid := child{5, "Chris", adult.Adult{}}
	fmt.Println(kid)
	kid.Adult.SetAge(25)
	fmt.Println(kid)
}