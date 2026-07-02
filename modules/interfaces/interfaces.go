package interfaces // interface is a reserved keyword in Go


import "fmt"

type Car string
type MotorCycle string

type Drive interface {
	Steer()
}

func (c Car) Steer(){
	fmt.Println("A car can steer!!")
}

func (m MotorCycle) Steer(){
	fmt.Println("A motorcycle can steer!!")
}

func canSteer(d Drive){ // using interface Drive
	d.Steer()
}

func LearnInterface(){

	canSteer(Car("Car"))
	canSteer(MotorCycle("MotorCycle"))	

}
