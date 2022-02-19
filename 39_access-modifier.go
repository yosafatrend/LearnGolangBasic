// public : diawali dengan huruf besar
// private : diawali dengan huruf kecil (camel case)
package main

import (
	"LearnGolangBasic/38_package-import/helper"
	"fmt"
)

func main() {
	helper.SayHello("Yosafat")
	//helper.sayGoodbye("Yosafat") //error

	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // error
}
