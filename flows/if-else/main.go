package main

import "fmt"

func main() {

	/*
		if statement - executa o conditie care respecta o anumita regula
		este una dintre cele mai simple structuri de control in Golang
		poate fi inlocuit uneori cu switch

	*/

	var x = 30
	if x%5 == 0 {
		fmt.Printf("%d este multiplu de 5\n", x)
	}

	// putem combina conditii multiple utilizand operatorii && si ||
	var varsta = 20
	if varsta >= 17 && varsta <= 25 {
		fmt.Println("Varsta este cuprinsa intre 17 si 25")
	}

	// exemplu if-else statement
	var varsta2 = 15
	if varsta2 >= 18 {
		fmt.Println("Poti da admitere la facultate!")
	} else {
		fmt.Println("Nu poti da admitere la facultate!")
	}

	// exemplu lant if-else-if
	var greutate = 78
	if greutate < 75 {
		fmt.Println("Sub greutatea normala")
	} else if greutate >= 75 && greutate <= 90 {
		fmt.Println("Greutate normala")
	} else if greutate > 90 && greutate <= 100 {
		fmt.Println("Greutate depasita")
	} else {
		fmt.Println("Probleme de sanatate")
	}

	/*
		if - short statement
		variabila va fi declarata si initializata in aceeasi structura
	*/
	if n := 20; n%2 == 0 {
		fmt.Printf("%d este par\n", n)
	}

	// variabila declarata in "short statement" poate fi folosita doar in blocul if-else
	if n := 15; n%2 == 0 {
		fmt.Printf("%d este par\n", n)
	} else {
		fmt.Printf("%d este impar\n", n)
	}
}
