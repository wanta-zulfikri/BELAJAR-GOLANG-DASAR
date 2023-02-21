package main

import "fmt"

func main () {
	name := "wanta"

	if name == "wanta"{
	   fmt.Println("hello wanta")
	} else if name == "joko" {
		fmt.Println("hello joko")
	} else if name == "budi" {
		fmt.Println("hello budi")
		fmt.Println("hI boleh kenalan ?")
	}


	
	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
