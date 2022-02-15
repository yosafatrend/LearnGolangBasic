package main

import "fmt"

func getFullName() (string, string, string) {
	return "Yosafat", "Rendi", "Prayoga"
}

func main() {
	firstname, middleName, _ := getFullName()

	fmt.Println(firstname)
	fmt.Println(middleName)
}
