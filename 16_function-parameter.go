package main

import "fmt"

func main() {
	sayHelloTo("Yosafat", "Rendi")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
