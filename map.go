package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Taufiq",
		"address": "Semarang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	//	make map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["auhor"] = "Taufiq"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
