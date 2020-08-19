package main

import "fmt"

func main() {

	/*
		panic si recover pot fi vazute la fel cu try-catch-finally din alte limbaje, precum Java
		panic si recover trebuie folosite in situatiile in care programul nu-si poate recupera executia
		cazurile valide pentru utilizarea panic sunt:
			o eroarea irecuperabila unde programul nu-si mai poate continua executia
			o eroarea a programatorului
		cand ajungem in cazul de panic, executia functiei e opreste, dar se vor executa functiile de tip defer
		functiile defer sunt functii care se apeleaza inainte ca functia din care sunt apelate sa isi incheie executia
	*/

	nume := "John"
	//prenume := "Smith"
	fullName(&nume, nil)
	fmt.Println("Executia din main pentru exemplul 1 s-a terminat cu succes")

	slicePanic()
	fmt.Println("Executia din main pentru exemplul slice s-a terminat cu succes")

}

// exemplu panic creat de programator
// pas 2 - adaiharea functie defer
func fullName(nume *string, prenume *string) {
	//defer fmt.Println("apelare defer in fullName")
	defer recoverFullName()
	if nume == nil {
		panic("runtime error: numele nu poate fi nil")
	}
	if prenume == nil {
		panic("runtime error: prenumele nu poate fi nil")
	}
	fmt.Printf("%s %s\n", *nume, *prenume)
	fmt.Println("Executia pentru exemplul 1 s-a terminat cu succes")
}

// exemplu de panic creat de o eroarea la runtime
func slicePanic() {
	n := []int{1, 2, 3}
	fmt.Println(n[4])
	fmt.Println("Executia in slice s-a terminat cu succes")
}

// recuperarea din panica
// dupa recuperare, panica se opreste si controlul revine apelantului
func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recuperare din ", r)
	}
}
