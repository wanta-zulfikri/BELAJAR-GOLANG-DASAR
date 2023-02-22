package main

import "fmt"

func factorialloop(value int) int {
	result := 1 
	for i := value; i >  0 ; i--{
		result *= i 
	
	}
	return result
} 

func factorialrecursive(value int) int {
	if value == 1 {
		return 1 
	} else {
		return value * factorialrecursive(value-1)
	}
}

func main () {
	loop := factorialloop(9)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialrecursive(5)
	fmt.Println(recursive)
}