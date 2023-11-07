package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // eko

	next := head.Next() // kurniawan
	fmt.Println(next.Value)

	next = next.Next() // khannedy
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
