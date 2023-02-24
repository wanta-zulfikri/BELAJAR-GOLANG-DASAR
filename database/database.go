package database

import "fmt"


var connection string

func init () {
	connection = "MYSQL"
}

func GetDataBase() string {
	return connection
}