package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke", counter)
	//	counter++
	//}
	//
	//fmt.Println("Selesai")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Selesai")

	names := []string{"eko", "kurniawan", "khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
