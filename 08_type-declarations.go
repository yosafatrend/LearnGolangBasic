package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "726481642844724"
	var marriedStatus Married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)

}
