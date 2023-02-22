package main 

import "fmt"


func main() {
	name := "wanta"
	counter := 0 
	increment := func() {
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}