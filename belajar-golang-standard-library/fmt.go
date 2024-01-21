package main

import "fmt"

func main() {
	firstName := "Eko"
	lastName := "Kurniawan"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
