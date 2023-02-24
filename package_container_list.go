package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("wanta")
	data.PushBack("zul")
	data.PushBack("fikri")
	data.PushFront("ka")

	for element := data.Front(); element == nil; element = element.Next(){
		fmt.Println(element.Value)
	}
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}