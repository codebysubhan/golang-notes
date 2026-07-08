package errorsingo


import "fmt"

func ErrorsInGo(){
	doSomething() // notice that its not like python and the order of the function doesn't matter
	fmt.Println("After doSomthing()")
}

func doSomething(){ // when this program runs the function "panics", and the defer is executed. But notice if I remove panic from here then the order of execution of defer is changed and when panic is called the order is different. what that means is that when some function is panicing then we can run a piece of code right before the panic actually gets executed. This comes in handy when we have to "recover" the function from somewhere.
	// defer fmt.Println("deferred") // we can recover using Recover()
	defer letsRecover() // now the function does panic anymore
	fmt.Println("1")
	panic(doSomething)
	fmt.Println("2")
}

func letsRecover(){
	recover()
}