package main

import "fmt"

func main() {
	/*
		in Golang un pointer este reprezentat de caracterul "*", urmat de tipul valorii
		* este, de asemenea, folosit pentru a dereferentia variabilele de tip pointer
		dereferentierea unui pointer ne ofera acces la valoarea stocata in pointerul respectiv
		*xPtr = 0 - stocheaza intregul 0 in zona de memorie referentiata de xPtr
		xPtr = 0 - eroarea de compilare pentru ca xPtr nu este un int, ci este un *int
		& este operatorul folosit pentru a gasi adresa unei variabile
		&x returneaza *int - pointer catre int, consideran ca este un intreg
	*/

	x := 5
	asignNoPointer(x)
	fmt.Println("Valoarea x este: ", x)

	xPtr := 5
	asignPointer(&xPtr)
	fmt.Println("Valoarea xPtr este: ", xPtr)

	// new primeste tipul ca si argument, aloca memorie si returneaza un pointer
	// in unele limbaje exista diferente mari intre new si &, dar Go suporta "garbage collection"
	xPtrNew := new(int)
	asignPointerNew(xPtrNew)
	fmt.Println("Valoarea xPtrNew este: ", *xPtrNew)

	xSquare := 1.5
	square(&xSquare)
	fmt.Println("Valoarea xSquare este: ", xSquare)
}

//argumentul este copiat in functie
//functia nu va modifica valoarea originala a variabilei x
func asignNoPointer(x int) {
	x = 3
}

//functia va modifica valoarea originala a variabilei xPtr
//refera o locatia din memorie
func asignPointer(xPtr *int) {
	*xPtr = 3
}

// modifica valoarea originala a variabilei xPtrNew
func asignPointerNew(xPtrNew *int) {
	*xPtrNew = 1
}

func square(x *float64) {
	*x = *x * *x
}
