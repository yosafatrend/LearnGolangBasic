package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yosafat",
		"address": "Pati",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "A Book"
	book["author"] = "Yosafat Rendi"
	book["other"] = "Nothing"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "other")

	fmt.Println(book)
	fmt.Println(len(book))

}
