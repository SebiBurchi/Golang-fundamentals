package main

import "fmt"

func main() {

	/*
		map este o colectie neordonata de tipul cheie valoare
		cheile sunt unice, dar valorile pot fi la fel
		este una dintre cele mai folosite colectii in programare
		valoarea zero pentru este nil
	*/

	// declarare map
	var m map[string]int // map cu cheile string si valorile de tip int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m este nil")
	}

	// initializare map utilizand functia make()
	fmt.Println("-----------------------Initializare map utilizand make si adaugare de element")
	var m1 = make(map[string]int)
	fmt.Println(m1)

	m1["str"] = 10 // adaugare element
	fmt.Println(m1)

	// initializare map utilizand map literal
	fmt.Println("-----------------------Initializare map utilizand map literal")
	var m2 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4, // virgula este necesara chiar daca este ultimul element
	}

	fmt.Println(m2)

	// preluarea unui element din map
	fmt.Println("-----------------------Preluare element din map")
	var m3 = map[string]int{
		"John":  1,
		"Steve": 2,
	}

	var johnValue = m3["John"]
	fmt.Println(johnValue)

	var steveValue = m3["Steve"]
	fmt.Println(steveValue)

	var noValue = m3["Ana"]
	fmt.Println(noValue) // pentru ca aceasta cheie nu exista, se va afisa valoarea zero pentru intregi

	// stergere cheie din map
	fmt.Println("-----------------------Stergere cheie din map")
	var fileExtensions = map[string]string{
		"Golang": ".go",
		"Python": ".py",
		"Java":   ".java",
	}
	fmt.Println(fileExtensions)

	delete(fileExtensions, "Python")
	fmt.Println(fileExtensions)

	// parcugere map
	fmt.Println("----------Parcugere map - Exemplu 1")
	var varstaMap = map[string]int{
		"John": 25,
		"Mike": 32,
		"Ana":  29,
	}

	for nume, varsta := range varstaMap {
		fmt.Println(nume, varsta)
	}

	fmt.Println("----------Parcugere map - Exemplu 1")
	for _, varsta := range varstaMap {
		fmt.Println(varsta)
	}

	// verificare existenta cheie in map
	fmt.Println("----------Verificare existenta cheie in map")
	salariiAngajati := map[string]int{
		"John": 12000,
		"Jack": 15000,
	}

	newAngajat := "John"
	value, ok := salariiAngajati[newAngajat]
	if ok == true {
		fmt.Println("Salariul pentru ", newAngajat, " este ", value)
		return
	}

	fmt.Println(newAngajat, "not found")

}
