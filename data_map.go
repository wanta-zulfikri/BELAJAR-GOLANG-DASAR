package main 

import "fmt"

func main () {
	person := map[string]string{
		"name" : "wanta",
		"address" : "tangerang",
	}
	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make (map[string]string)
	book["title"] = "belajar go-lang"
	book["author"] = "wanta"
	book["ups"]    = "salah"
	
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}