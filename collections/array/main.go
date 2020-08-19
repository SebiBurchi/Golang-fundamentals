package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		array - colectie alcatuita din elemente de un singur tip
		tine un anumit numar de elemente si nu poate creste in mod dinamic
		indexul primului element este 0
	*/

	// declararea unui array
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())

	// asignarea si accesarea elementelor in array
	var countryArray [3]string
	countryArray[0] = "Romania"
	countryArray[1] = "Spania"
	countryArray[2] = "Italia"

	fmt.Println("--------------------Asignarea si accesarea elementelor in array")
	fmt.Println(countryArray[0])
	fmt.Println(countryArray[1])
	fmt.Println(countryArray[2])

	// initializarea unui array cu valori predefinite - array literal
	x := [5]int{10, 20, 30, 40, 50}   // initializare cu valori
	var y [5]int = [5]int{10, 20, 30} // asignarea partiala

	fmt.Println("--------------------Array literal")
	fmt.Println(x)
	fmt.Println(y)

	/* initializarea unui array folosind ... in loc de specificarea dimensiunii
	compilatorul poate identifica dimensiunea unui array bazandu-se pe elementele specificate
	*/

	t := [...]int{10, 20, 30}

	fmt.Println("--------------------Initializare array folosind ...")
	fmt.Println(reflect.ValueOf(t).Kind())
	fmt.Println(len(t))

	// initializare array folosind array literal pentru a specifica anumite valori
	x1 := [5]int{1: 10, 3: 30}
	fmt.Println("--------------------Specificarea elementelor in array")
	fmt.Println(x1)

	// parcugere array folosind for loop
	testArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n---------Parcurgere array - exemplul 1")
	for i := 0; i < len(testArray); i++ {
		fmt.Println(testArray[i])
	}

	fmt.Println("\n---------Parcurgere array - exemplul 2")
	for index, element := range testArray {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("\n---------Parcurgere array - exemplul 3")
	for _, value := range testArray {
		fmt.Println(value)
	}

	// copierea unui array in alt array prin valoare sau referinta
	oldArray := [3]string{"Romania", "Germania", "USA"}
	newArray1 := oldArray  // datele sunt trimise prin valoare
	newArray2 := &oldArray // datele sunt trimise prin referinta

	fmt.Printf("oldArray: %v\n", oldArray)
	fmt.Printf("newArray1: %v\n", newArray1)

	oldArray[0] = "Canada"

	fmt.Printf("oldArray: %v\n", oldArray)
	fmt.Printf("newArray1: %v\n", newArray1)
	fmt.Printf("newArray2: %v\n", *newArray2)

	// verificarea existentei unui element in array
	checkArray := [5]string{"Romania", "Germania", "USA"}
	fmt.Println(itemExists(checkArray, "Romania"))
	fmt.Println(itemExists(checkArray, "Italia"))

}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Tip de date invalid")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
