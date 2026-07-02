package adult


import (
	"errors"
)

type Adult struct{ // to export a field we have to do capitalize the first letter
	Age int
	Name string
}

func (a *Adult) SetAge(newAge int) error{
	if newAge < 1{
		return errors.New("Invalid age")
	}
	a.Age = newAge
	return nil
}

func (a *Adult) GetAge() int{ // we still passed the pointer because Go's convention says that if we pass a pointer for one method then we should provide the pointer for all other methods as well
	return a.Age
}
