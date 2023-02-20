package main

import "fmt"

func main() {
	var nilai32 int
	var nilai64 int
	var nilai18 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
    fmt.Println(nilai18)

	var name = "wanta"
	var e byte = name[0]
	var eString string = string (e)

	fmt.Println(name)
	fmt.Println(eString)
}