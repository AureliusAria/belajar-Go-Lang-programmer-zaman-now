package main

import "fmt"

func main() {
	var name [4]string

	name[0] = "Aurelius"
	name[1] = "Aria"
	name[2] = "Baras"
	name[3] = "Panyapa"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	fmt.Println(name[3])

	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("Panjang data", len(values))
}
