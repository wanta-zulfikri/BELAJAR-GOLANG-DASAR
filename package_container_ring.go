package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	for i := 0; i < data.Len(); i ++ {
		data.Value="value" + strconv.FormatInt(int64(i), 10)
	}

	data.Do(func(value interface{}){
		fmt.Println(value)
	})

	var data1 *ring.Ring = ring.New(5)

	for i := 0; i < data1.Len() ; i ++ {
		data1.Value = "data " + strconv.FormatInt(int64(i), 10)

	}

	data1.Do(func(value interface{}) {
		fmt.Println(value)
	})
}