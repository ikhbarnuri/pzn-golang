package main

import "fmt"

func sayHello(firstName string, lastName string) string {
	return "Hello" + " " + firstName + " " + lastName
}

func main() {
	fmt.Println(sayHello("Eko", "Kurniawan"))
}
