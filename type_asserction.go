package main 

import "fmt" 
func random () interface {} {
	return "ups"
}

func main() {
	result := random() 
	resultString := result.(string)
	fmt.Println(resultString)

/**
	result := random()
	switch value := result.(type) {
	case int:
		fmt.Println("value", vslue, "is int")
	case string :
		fmt.Println("VALUE", value, "is string")
	default:
		fmt.Println("Unknown type")

	}
}
