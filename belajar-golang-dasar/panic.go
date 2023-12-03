package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups error")
	}
}

func main() {
	runApp(false)
	fmt.Println("Eko Kurniawan")
}
