package main

import "fmt"

func registerUser(name string, filter func(string) bool) {
	if filter(name) {
		fmt.Println("You're blocked " + name)
	} else {
		fmt.Println("Welcome " + name)
	}
}

//without anonymous function

// func blacklistAdmin(name string) bool{
// 	return name == "Admin"
// }

// func blacklistRoot(name string) bool{
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("Yosafat", blacklist)
	registerUser("admin", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}
