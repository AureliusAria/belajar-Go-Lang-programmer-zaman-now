package main

import "fmt"

func getSayHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}

}
func main() {
	result := getSayHello("Aria")
	fmt.Println(result)

	fmt.Println(getSayHello(""))

}
