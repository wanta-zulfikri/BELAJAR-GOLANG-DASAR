package main

import (
	"fmt"
	"reflect"
)

type sample struct {
	Name string 
}

type Contohlagi struct {
	Name string`required:"true`
	Description string
}

func Isvalid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i ++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false 
			}
		}
	}
	return true 
}

func main() {
	sample := sample{"eko"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(required)

}

type sample1 struct {
	Name string `required: "true" max:"10"`
}


	

	