package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("U are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "anjing"
	}
	registerUser("anjing", blackList)

	registerUser("Eko", func(name string) bool {
		return name == "anjing"
	})
}
