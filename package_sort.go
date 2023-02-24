package main 

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int 

}

type UserSlice []User

func (value UserSlice) len() int{
		return len(value)
}
func (value UserSlice) less(i, j int) bool {
	return value[i].Age < value [j].Age

}

func (value UserSlice) swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User {
		{"wanta", 30},
		{"eko", 33},
		{"ari", 39},
		{"dewq", 31},
		{"wa", 37},

	}
	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}