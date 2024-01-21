package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")
	data.PushFront(90)

	head := data.Front()
	fmt.Println(head.Value)

	head = head.Next()
	fmt.Println(head.Value)

	head = head.Next()
	fmt.Println(head.Value)

	head = head.Next()
	fmt.Println(head.Value)

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
