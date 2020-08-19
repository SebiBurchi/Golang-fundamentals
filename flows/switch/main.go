package main

import "fmt"

func main() {

	/*
		switch - potriveste o expresie cu o lista de cazuri posibile
		este folosit uneori pentru a inlocui multiple structuri if-else
		este diferit de alte limbaje cum ar fi Java, unde pentru a iesi dintr-un block switch este nevoie de break
	*/

	fmt.Println("\nExemplu switch")
	var zi = 6
	switch zi {
	case 1:
		fmt.Println("Luni")
	case 2:
		fmt.Println("Marti")
	case 3:
		fmt.Println("Miercuri")
	case 4:
		fmt.Println("Joi")
	case 5:
		fmt.Println("Vineri")
	case 6:
		{
			fmt.Println("Sambata")
			fmt.Println("Nu muncim")
		}
	case 7:
		fmt.Println("Duminica")
	}

	/*
		switch - short statement
		diferenta este ca variabila folosita este valabila doar in interiorul structurii de control switch
	*/

	fmt.Println("\nShort statement switch")
	switch zi := 8; zi {
	case 1:
		fmt.Println("Luni")
	case 2:
		fmt.Println("Marti")
	case 3:
		fmt.Println("Miercuri")
	case 4:
		fmt.Println("Joi")
	case 5:
		fmt.Println("Vineri")
	case 6:
		{
			fmt.Println("Sambata")
			fmt.Println("Nu muncim")
		}
	case 7:
		fmt.Println("Duminica")
	default:
		fmt.Println("Gresit")
	}

	/*
		switch fara expresie
		in Golang, expresia specificata in switch este optionala
		un switch fara expresie este echivalentul unui "switch true"
		evalueaza toate cazurile si ruleaza primul caz care este adevarat
		switch fara expresie este o alternativa mai concisa de a scris lantul if-else-if
	*/

	fmt.Println("\nSwitch fara expresie")
	var greutate = 22.0
	switch {
	case greutate < 18.5:
		fmt.Println("Sub greutatea normala")
	case greutate >= 18.5 && greutate < 25.0:
		fmt.Println("Greutatea este normala")
	case greutate >= 25.0 && greutate < 30.0:
		fmt.Println("Greutate marita")
	default:
		fmt.Println("Probleme de sanatate")
	}

	// combinarea de cazuri multiple in switch
	fmt.Println("\nCombinarea de cazuri multiple in switch")
	switch zi := 6; zi {
	case 1, 2, 3, 4, 5:
		fmt.Println("Zi de lucru")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Gresit")
	}
}
