package main

import "fmt"

func main() {

	var ujian = 80
	var absen = 75

	var lulusUjian = ujian >= 80
	var lulusAbsen = absen >= 80

	var lulus = lulusUjian && lulusAbsen

	fmt.Println(lulus)
}
