package main 

import "fmt"


type filter func (string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("wanta", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}
