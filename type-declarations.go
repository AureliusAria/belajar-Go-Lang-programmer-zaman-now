package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var noKtpAria NoKTP = "1232232"
	var marriedStatus Married = false
	fmt.Println(noKtpAria)
	fmt.Println(marriedStatus)

}
