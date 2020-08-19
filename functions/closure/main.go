package main

import "fmt"

func main() {

	/*
		functii anonime
		o functie anonima este la fel ca o functie normala, doar ca nu are un nume
		fara un nume, o functie anonima este creata dinamic, asemanator unei variabile
	*/

	//apelarea functiei anonime - avantaj: putem schimba in mod dinamic comportamentul
	fmt.Println("---------Apelare functie anonima")
	DoStuff()

	DoStuff = func() {
		fmt.Println("Fa ceva!")
	}
	DoStuff()

	DoStuff = func() {
		fmt.Println("Fa altceva!")
	}
	DoStuff()

	/*
		closure este un tip special de functie anonima care referentiaza o variabila declarata in afara functiei in sine
		acest comportament este foarte similar cu modul in care o functie normala poate referentia variabile globale
	*/

	// exemplu closure
	fmt.Println("------Exemplu closure")
	n := 0
	counter := func() int {
		n++
		return n
	}
	fmt.Println(counter())
	fmt.Println(counter())

	// functia recursiva se apeleaza pe ea

	rezultat := recursiva(4)
	fmt.Printf("Recursiva: %d\n", rezultat)
}

// declararea unei functii anonime
var DoStuff func() = func() {
	// Fa ceva
}

// declararea unei functii recursive
func recursiva(number int) int {
	if number == 1 {
		return number
	}

	return number + recursiva(number-1)
}
