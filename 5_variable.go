package main

import "fmt"

func main() {
	var name string

	name = "Yosafat Rendi"
	fmt.Println(name)

	name = "Rendi Prayoga"
	fmt.Println(name)

	var friendName = "Akanichiyu" //without data type
	fmt.Println(friendName)

	var age = 47
	fmt.Println(age)

	country := "Indonesia" //without "var"
	fmt.Println(country)

	var (
		firstName = "Yosafat"
		lastName  = "Rendi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
