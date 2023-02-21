package main

import "fmt" 

func sumALL(numbers ...int) int {
	total := 0 
	for _, number := range numbers {
		total += number 
	}
	return total 
}

func main() {
	total := sumALL(10, 90, 30, 50, 40)
	fmt.Println(total)
} 

// slice parameter 
func main () {
	total := sumALL(10, 90, 30, 50, 40)
	total = sumALL(slice...)
	