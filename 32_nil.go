// nil sama lah kayak null, tapi ini untuk beberapa tipe data aja
package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var data map[string]string = NewMap("Yosafat")

	if data == nil {
		fmt.Println("Data Kosong!")
	} else {
		fmt.Println(data)
	}
}
