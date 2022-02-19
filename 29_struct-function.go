package main

import "fmt"

type People struct {
	Name, Address string
	Age           int
}

func (people People) sayHi(name string) {
	fmt.Println("Hi", name, "my name is", people.Name)
}

func main() {
	people := People{
		Name:    "Yosafat",
		Address: "Pati",
		Age:     19,
	}

	fmt.Println(people)
	people.sayHi("Akani")
}
