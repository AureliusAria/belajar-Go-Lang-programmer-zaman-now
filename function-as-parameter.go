package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}
func spamFilter(name string) string {
	if name == "Anjing" || name == "anjing" || name == "ajg" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Aria", spamFilter)
	sayHelloWithFilter("ajg", spamFilter)

}
