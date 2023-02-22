package main 

import "fmt"

func main () {
	var month = [12] string {
		"january",
		"febuary",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"october",
		"november",
		"desember",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2,"wanta")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"

	fmt.Println(slice2)
	fmt.Println(month)

	newslice := make ([]string, 3, 5)

	newslice[0] = "wanta"
	newslice[1] = "zulfikri"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copyslice := make([]string, len(newslice), cap(newslice))
	copy(copyslice, newslice)
	fmt.Println(copyslice)

	iniarray := [5]int{1,2,3,4,5}
	inislace := []int {1,2,3,4,5}

	fmt.Println(iniarray)
	fmt.Println(inislace)
}