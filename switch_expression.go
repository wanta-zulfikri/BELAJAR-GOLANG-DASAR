package main

import "fmt"

func main () {
	name := "rian"

	switch name {
	case "wanta" :
		fmt.Println("hello wanta")
    case "joko"  :
		fmt.Println("HELLO JOKO")
	case "rian" :
		fmt.Println("HELLO RIAN")
	}

	
	switch length := len(name); length > 5 {
	case true :
		fmt.Println("nama terlalu panjang")
	case false :
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
		
	
}