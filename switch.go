package main

import "fmt"

func main() {

	name := "Aria"

	switch name {
	case "Aria":
		fmt.Println("Hello Aria")
		fmt.Println("Hello Aria")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	case "Budi":
		fmt.Println("Hello Budi")
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Kenalan dong!")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}


	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
