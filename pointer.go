package main 

import "fmt"

type Address struct {
	city, province, country string
}

func main () {
	var address1 Address = Address{"subang", "jawa barat", "indonesia"}
	var address2 *Address = &address1 
	var address3 *Address = &address1
	address2.city = "bandung"

	*address2 = Address{"jakarta", "Dki Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.city = "jakarta"
	fmt.Println(address4)
}