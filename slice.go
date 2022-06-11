package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Mei lagi"
	fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Taufiq")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Taufiq"
	newSlice[1] = "Kurniawan"
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//	carefully about array and slice

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{12, 3, 4, 5,}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
