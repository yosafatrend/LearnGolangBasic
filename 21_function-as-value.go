package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	sayGoodbye := getGoodbye

	result := sayGoodbye
	fmt.Println(result("Yosafat"))
	//same as
	fmt.Println(getGoodbye("Yosafat"))
}
