package main 

import "fmt" 

func main () {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
	// for dengan statement 
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}
    // for range 
	slice := []string{"wanta", "zul", "fikri"}

	for i := 0; i < len (slice); i++ {
		fmt.Println(slice[i])

	}


	// for range 
	for i, value := range slice{
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "wanta"
	person["title"] = "progremmer"
}