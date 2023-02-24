package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("100000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
	value := strconv.FormatInt(100000, 2)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("20000000")
	fmt.Println(valueInt)
}
