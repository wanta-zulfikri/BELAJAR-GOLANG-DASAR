package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) maried() {
	man.Name = "mr. " + man.Name
	fmt.Println("di method", man.Name)

}

func main() {
	eko := Man{"eko"}
	eko.maried()

	fmt.Println(eko.Name)
}
