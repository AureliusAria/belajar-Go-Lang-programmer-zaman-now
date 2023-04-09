package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan", counter)
	}

	fmt.Println("\n", "for")

	slice := []string{"Aurelius", "Aria", "Baras", "Panyapa"}
	//fmt.Println(slice[2])

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("\n", "for range")
	for _, name := range slice {
		//fmt.Println("index ", i, " = ", name)

		//jika variabel tidak terpakai maka bisa diganti _ supaya tidak error
		fmt.Println(name)
	}

	fmt.Println("\n", "Map")

	person := make(map[string]string)
	person["Name"] = "Aria"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}

}
