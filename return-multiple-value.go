package main

import "fmt"

func getFullNmae() (string, string, string) {
	return "Aurelius", "Aria", "Baras"

}

func main() {
	firstName, middleName, _ := getFullNmae()
	fmt.Println(firstName)
	fmt.Println(middleName)
	//fmt.Println(lastName)

	fmt.Println(getFullNmae())
}
