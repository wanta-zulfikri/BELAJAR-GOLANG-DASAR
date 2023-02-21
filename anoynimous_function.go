package main

import "fmt"

type blacklist func(string) bool 

func registeruser(name string, blacklist blacklist) {
     if blacklist(name){
		fmt.Println("you are blocked", name)
	 } else {
		fmt.Println("welcome", name)
	 }
	 
}


func main () {
	blacklist := func (name string ) bool {
		return name == "admin"
	}

	registeruser("admin", blacklist)
	registeruser("wanta", blacklist)

	registeruser("root", func(name string) bool{
		return name == "root"
	})
	registeruser("wanta", func(name string) bool {
		return name== "root"
	})
}