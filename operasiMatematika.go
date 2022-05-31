package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 20
		c     = a + b

		i int = 20
	)
	i += 20
	i++
	i--
	fmt.Println(c)
	fmt.Println(i)

}
