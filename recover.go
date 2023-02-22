package main

import "fmt"
func endapp(){
	message := recover()
	fmt.Println("Error dengan message:", message)
	fmt.Println("Aplikasi selesai")
}

func runapp(error bool) {
	defer endapp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("APLIKASI BERJALAN")
}
func main () {
	runapp(true)
}