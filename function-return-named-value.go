package main

import "fmt"

func getFullName2() (firstName, middleName, lastname string) {
	firstName = "Aurelius"
	middleName = "Aria"
	lastname = "Baras"

	return
}

func main() {
	a, b, c := getFullName2()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	getFullName2()
}
