package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello wanta"
	} else {
	return "Hello" + name 
	}

}



func main () {
	result := getHello("wanta")
	fmt.Println(result)
	
	fmt.Println(getHello(""))

}