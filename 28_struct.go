package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var yosafat Customer
	yosafat.Name = "Yosafat"
	yosafat.Address = "Pati, Jawa Tengah"
	yosafat.Age = 19

	fmt.Println(yosafat.Name)
	fmt.Println(yosafat.Address)
	fmt.Println(yosafat.Age)

	aka := Customer{
		Name:    "Akanichiyu",
		Address: "Salatiga, Jawa Tengah",
		Age:     18,
	}
	fmt.Println(aka)

	rodori := Customer{"Rodorimishi", "Salatiga, Jawa Tengah", 15}
	fmt.Println(rodori)
}
