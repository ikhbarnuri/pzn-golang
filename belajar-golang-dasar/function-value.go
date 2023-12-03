package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodbye
	fmt.Println(contoh("Eko"))
}
