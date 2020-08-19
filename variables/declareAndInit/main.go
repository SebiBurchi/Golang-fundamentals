package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	/*
		variabila reprezinta numele dat unei locatii in memorie pentru a stoca o valoare de un anumit tip
	*/

	/* declararea unei variabile
	sintaxa: var name type
	*/

	var varsta int                       // declararea variabilei
	fmt.Println("Varsta este: ", varsta) //pentru ca nu am asignat nici o valoare, valoarea default este 0

	// odata declarata, variabila poate fi asignata catre orice valoare din acel tip
	varsta = 30
	fmt.Println("Varsta este: ", varsta)

	/*
		declararea unei variabile cu valoare initiala
		sintaxa: var name type = initialValue
	*/
	var varsta2 int = 22 //declarare + initializare variabila
	fmt.Println("Varsta2 este: ", varsta2)

	/*
		type inference
		daca variabila are o valoare initiala, Go va face type inference pentru a obtine tipul variabilei
	*/
	var varsta3 = 29 //declararea variabilei fara tip -> type inference
	fmt.Println("Varsta3 este: ", varsta3)
	fmt.Println("Varsta3 este de tipul: ", reflect.TypeOf(varsta3))

	/* declarari multiple de variabile
	sintaxa: var name1, name2 type = initialValue1, initialValue2
	*/
	var lungime, latime int = 100, 50 //declarare + initializare variabile multiple
	fmt.Println("Lungime este: ", lungime, " latimea este: ", latime)

	//declarari si initializari multiple + type inference
	var lungime2, latime2 = 100, 50
	fmt.Println("Lungime2 este: ", lungime2, " latimea 2 este: ", latime2)

	/*
		declararea si initializarea de variabile de tipuri diferite in acelasi bloc
		sintaxa: var (
			name1 = initialValue1
			name2 = initialValue2
			......
		)
	*/
	var (
		nume     = "telacad"
		varsta4  = 22
		inaltime int
	)
	fmt.Println("Numele este: ", nume, ", varsta este: ", varsta4, ", inaltimea este: ", inaltime)

	/*
		declararea si initializarea de variabile folosind conceptul de "short hand declaration" (operatorul :=)
		sintaxa: name := initialValue
		poate fi folosit cand cel putin una din variabile este noua
	*/
	count := 20
	fmt.Println("Count=", count)

	nume2, varsta5 := "telacad", 40
	fmt.Println("numele este: ", nume2, "varsta este: ", varsta5) // "short hand declaration" pentru variabile multiple

	// declararea si initializarea unei variabile la runtime
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Valoarea minima este: ", c)

	// strongly typed -> variabilele declarate ca apartinand unui tip de date nu pot fi asignate catre alt tip
	varsta6 := 33
	//varsta6 = "telacad"
	fmt.Println(varsta6)

	/*
		constantele nu-si mai pot schimba valoarea dupa ce au fost initializate
		sunt definite utilizand cuvantul cheie const
		numele este scris, de obicei, cu litere mari
	*/
	const PRET = 100
	fmt.Println("pretul este: ", PRET)

	// PRET = 101  // eroarea de compilarea deoarece PRET este o constanta

}
