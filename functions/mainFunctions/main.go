package main

import "fmt"

func main() {

	/*
		functia este un grup de declaratii
		exista in cadrul unui program cu scopul indeplinirii unei sarcini specifice
		la nivel inalt, o functie preia o intrare si returneaza o iesire
		cea mai cunoscuta functie in Go este main, functia prinicipala
		pentru definirea unei functii se foloseste cuvantul cheie func

	*/

	// apelarea functie simple
	fmt.Println("---------Apelare functie simpla")
	firstFunction()

	// apelarea functie cu parametri/argumente
	fmt.Println("---------Apelare functie cu parametri")
	total(20, 30)

	// apelare functie cu argumente si tip returnat
	fmt.Println("---------Apelare functie cu argumente si tip returnat")
	sum := totalReturn(20, 30) // preluarea valoare returnata intr-o variabila
	fmt.Println(sum)

	// apelarea functie cu tip returnat numit
	fmt.Println("---------Apelare functie cu tip returnat numit")
	fmt.Println("Aria: ", dreptunghi(20, 30))

	// apelarea functie cu valori multiple returnate
	fmt.Println("---------Apelare functie cu valori multiple returnate")
	var a, p int
	a, p = dreptunghi2(20, 30)
	fmt.Println("Aria: ", a)
	fmt.Println("Perimetru: ", p)

	// utilizat blank identifier (_)
	var p2 int
	_, p2 = dreptunghi2(20, 30)
	fmt.Println("Perimetru2: ", p2)
}

// creare unei functii simple
func firstFunction() {
	fmt.Println("Prima functie")
}

// crearea unei functii cu parametri/argumente
func total(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// crearea unei functii cu tip returnat numit
func dreptunghi(lungime int, latime int) (aria int) {
	var perimetru int
	perimetru = 2 * (lungime + latime)
	fmt.Println("Perimetrul: ", perimetru)

	aria = lungime * latime
	return // returnam fara a specifica valoarea returnata
}

// crearea unei functii cu tip returnat
func totalReturn(x int, y int) int {
	total := 0
	total = x + y
	return total
}

// crearea unei functii cu valori multiple returnate
func dreptunghi2(lungime int, latime int) (aria int, perimetru int) {
	perimetru = 2 * (lungime + latime)
	aria = lungime * latime
	return // fara a specifica valorile returnate

}
