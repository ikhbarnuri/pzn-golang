package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("EKo")
	fmt.Println(result)
	fmt.Println(helper.version)
	fmt.Println(helper.Application)
	fmt.Println(helper.sayGoodbye("Budi"))
}
