package main

import "fmt"

func main() {
	counter := 1
	for counter <= 5 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	slice := []string{"Yosafat", "Rendi", "Prayoga"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := map[string]string{
		"name":  "Yosafat Rendi",
		"title": "Programmer",
	}

	for key, value := range person {
		fmt.Println(key, value)
	}
}
