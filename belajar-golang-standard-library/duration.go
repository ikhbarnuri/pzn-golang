package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Second * 1000
	duration2 := time.Millisecond * 10

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Printf("duration: %d\n", duration2)
}
