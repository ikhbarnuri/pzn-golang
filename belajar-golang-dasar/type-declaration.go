package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "11111111111"

	var contoh string = "222222222222"
	var contohKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKTP)
}
