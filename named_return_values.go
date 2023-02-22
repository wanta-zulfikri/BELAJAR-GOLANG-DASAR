package main 

import "fmt" 

func getFullname2() (fisrtname string, middlename string, lastname string) {
	fisrtname = "wanta"
	middlename = "zul"
	lastname = "fikri"
	return
}
func main () {
	a, b, c := getFullname2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}