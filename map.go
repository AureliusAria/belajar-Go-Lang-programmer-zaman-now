package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Aria",
		"address": "Pontianak",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book = make(map[string]string)

	book["title"] = "Belajar-go-dasar"
	book["author"] = "Aria"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
