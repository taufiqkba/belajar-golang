package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var NoKu NoKTP = "332191921921"
	var marriedStatus Married = true

	fmt.Println(NoKu)
	fmt.Println(marriedStatus)
}
