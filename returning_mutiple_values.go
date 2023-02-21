package main 

import "fmt" 

func getFullname() (string, string) {
	return "wanta", "zulfikri"
}

func main () {
	fisrtname, lastname := getFullname()
	fmt.Println(fisrtname, lastname)
    fmt.Printf(lastname)
} 