package main

import "fmt"

type Person struct {
	FirstName string
	lastName  string
	Age       int
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{firstname, lastname, age}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{firstname, lastname, age}
}


func main() {
	sl := make([]Person,0)
	fmt.Println("before")
	for i:=0; i< 10_000_000;i++{
		sl = append(sl, MakePerson("apalah","injuga",34))
	}

	fmt.Println("after", len(sl))
}
