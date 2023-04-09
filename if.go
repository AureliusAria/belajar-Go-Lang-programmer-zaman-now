package main

import "fmt"

func main() {

	var name = "Rinaaaa"

	if name == "Aria" || name == "aria" || name == "ARIA" {
		fmt.Println("Hello Aria")
	} else if name == "Joko" || name == "joko" || name == "JOKO" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" || name == "budi" || name == "BUDI" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hai, Kenalan dong!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
