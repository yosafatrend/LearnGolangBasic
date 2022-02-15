package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("(b)Perulangan ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("(c)Perulangan ke", i)
	}
}
