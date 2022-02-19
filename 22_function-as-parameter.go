package main

import "fmt"

// efficiency for long func, use type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello " + filteredName)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Yosafat", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
