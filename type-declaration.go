package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpTaufiq NoKTP = "0219392"
	var marriedStatus Married = true
	fmt.Println(noKtpTaufiq)
	fmt.Println(marriedStatus)
}
