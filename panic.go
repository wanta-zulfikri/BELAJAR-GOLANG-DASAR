package main 

import "fmt"

func endapp() {
	fmt.Println("aplikasi selesai")
}

func runapp(error bool){
	defer endapp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main () {
	runapp(false)
}