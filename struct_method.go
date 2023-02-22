package main 

import "fmt"

type Customer struct {
	Name, address string 
	Age 		int 
}

func (Customer Customer) sayHello3() {
	fmt.Println("Hello , my name is", Customer.Name)

}

func main () {
	var eko Customer
	eko.Name = "eko"
	eko.address = "indonesia"
	eko.Age = 30

	eko.sayHello3("Rian")
	eko.sayHello3()
}