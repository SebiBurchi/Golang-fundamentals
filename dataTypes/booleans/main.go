package main

import "fmt"

func main() {

	/*
		Booleans
		Tipul de data in Golang este bool
		Valorile pot fi true sau false
	*/

	t := true
	f := false
	fmt.Println("t != f: ", t != f)

	//valoarea default pentru bool este false
	var booleanType bool
	fmt.Printf("%v\n", booleanType)

	//rezultatul compararii unor valori este un boolean
	a := 5
	b := 8
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
}
