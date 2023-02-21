package main

import "fmt"


type Customer struct {
	Name, address string 
	Age 		int 
}

func main() {


// membuat data struct 

	var eko Customer
	eko. Name = "eko"
	eko.address   = "indonesia"
	eko.Age = 40

	fmt.Println(eko.Name)
	fmt.Println(eko.address)
	fmt.Println(eko.Age)

	// struct literals
	joko := Customer {
		Name : "joko",
		address:  "indonesia",
		Age : 30 ,
	}
	fmt.Println(joko)

	budi := Customer{"budi","jakarta",35}
	fmt.Println(budi)
}