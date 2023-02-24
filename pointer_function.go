package main 

import "fmt"


func changeaddresstoindonesia(address *Address) {
	address.country = "indonesia"
} 
      var alamat = Address{
		city:  "subang",
		province : "jawa barat",
		country: "" ,
	  }


func main() { 
	address := Address{"subang", "jawa", ""}
	changeaddresstoindonesia(&address)

	fmt.Println(address) // berubah
}