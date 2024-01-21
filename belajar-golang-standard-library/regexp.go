package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex, _ := regexp.Compile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eeko"))
	fmt.Println(regex.MatchString("ezo"))

	fmt.Println(regex.FindAllString("eko eeko ezo ego eKo", 10))
}
