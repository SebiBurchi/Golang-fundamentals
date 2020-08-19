package main

import "fmt"

func main() {

	/*
		goto statement
		este folosit pentru a altera executia normala a unui program
		transfera controlul catre o eticheta (label) in acelasi program
		eticheta este un identificator, poate text chior si poate fi setata oriunde in programul Go
		transfera controlul catre label si incepe executia de acolo
	*/

	var varsta int
votare:
	fmt.Printf("Introdu varsta:")
	fmt.Scanln(&varsta)
	if varsta <= 17 {
		fmt.Println("Nu esti eligibil pentru a vota")
		goto votare
	} else {
		fmt.Println("Esti eligibil pentru a vota")
	}

}
