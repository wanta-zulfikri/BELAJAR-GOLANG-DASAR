package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("wanta zulfikri", "wanta"))
	fmt.Println(strings.Split("wanta zulfikri", " "))
	fmt.Println(strings.ToLower("wanta zulfikri"))
	fmt.Println(strings.ToUpper("wanta zulfikri"))
	fmt.Println(strings.Trim("   wanta zulfikri   ", " "))
	fmt.Println(strings.ReplaceAll("wanta wanta wanta wanta", "wanta", "fikri"))

}