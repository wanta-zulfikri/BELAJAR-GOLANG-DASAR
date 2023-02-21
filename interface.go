package main 

import "fmt"

type HasName interface {
	GetName() string 
}


func sayHello4(hasName HasName) {
	fmt.Println("HELLO", hasName.GetName())
}
func sayHello3(Animal Animal) {
	fmt.Println("HELLO", Animal.GetName())
}

type person struct {
	Name string 
}
type Animal struct {
	Name string 
}
func (person person) GetName() string {
	return person.Name
func (Animal Animal) GetName() string {
	return Animal.Name
}
func main() {
	var eko person
	eko.Name = "eko"

	sayHello4(eko)

	cat := Animal{
		Name : "kuciang",
	}
	sayHello3(cat)
}