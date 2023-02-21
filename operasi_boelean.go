package main

import "fmt"

func main (){
	var nilaiakhir = 90 
	var absensi = 80 

	var lulusnilaiakhir bool = nilaiakhir >= 80 
	var lulusanabsensi bool = absensi >= 80 
	fmt.Println(lulusanabsensi)
	fmt.Println(lulusnilaiakhir)

	var lulus = lulusanabsensi && lulusnilaiakhir
	fmt.Println(lulus)

	fmt.Println(nilaiakhir >= 80 && absensi >= 80)
}