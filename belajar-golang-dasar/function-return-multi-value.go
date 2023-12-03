package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
