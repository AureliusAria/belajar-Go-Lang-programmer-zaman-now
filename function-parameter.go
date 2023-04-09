package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
func main() {
	firstName := "Aurelius"
	sayHello(firstName, "Aria")
	sayHello("Budi", "Nugraha")

}
