package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	values := []int{100, 90, 80}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "Eko"))
	fmt.Println(slices.Index(names, "Ekos"))
}
