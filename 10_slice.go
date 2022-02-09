package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Changed"
	// fmt.Println(slice1)

	// slice1[0] = "AnotherMay"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Yosafat") //if capacity of array reach the limit while append, it will make a new array so the real array won't changed
	fmt.Println(slice3)
	slice3[1] = "NotDecember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yosafat"
	newSlice[1] = "Rendi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisIsArray := [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
