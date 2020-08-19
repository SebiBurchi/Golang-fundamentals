package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		slice - colectie flexibila, dar si extensibila
		este un segment de tablouri dinamice care pot creste si micsora dupa cum doreste utilizatorul
		este de cele mai multe ori preferat in detrimentul array-ului
	*/

	// crearea unui slice empty
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	// crearea unui slice utilizand functia make
	var intSlice1 = make([]int, 10)        // atunci cand dimensiunea si capacitatea slice-ului sunt la fel
	var strSlice1 = make([]string, 10, 20) // dimensiunea si capacitatea difera

	fmt.Println("\n----------------------Creare slice folosind make")
	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	fmt.Println(reflect.ValueOf(intSlice1).Kind())

	fmt.Printf("strSlice1 \tLen: %v \tCap: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Println(reflect.ValueOf(intSlice1).Kind())

	// initializarea unui slice folosind valori - slice literal
	var intSlice2 = []int{10, 20, 30, 40}
	var strSlice2 = []string{"John", "Ana", "Gigi"}

	fmt.Println("\n------------------------Initializare slice folosind valori - slice literal")
	fmt.Printf("intSlice2 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
	fmt.Printf("strSlice2 \tLen: %v \tCap: %v\n", len(strSlice2), cap(strSlice2))

	// crearea slice utilizand cuvantul cheie new
	var newSlice = new([50]int)[0:10] // declararea capacitatii (50), urmata de dimensiune (10)

	fmt.Println("\n------------------------Creare slice utilizand new")
	fmt.Printf("newSlice \tLen: %v \tCap: %v\n", len(newSlice), cap(newSlice))
	fmt.Println(newSlice)

	// adaugarea de elemente in slice utilizand functia append()
	fmt.Println("\n------------------------Adaugare de elemente in slice")
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice a: ", a)
	fmt.Printf("Dimensiunea este %d, Capacitatea este %d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice a dupa adaugarea de elemente: ", a)
	fmt.Printf("Dimensiunea este %d, Capacitatea este %d\n", len(a), cap(a))

	// accesarea elementelor din slice
	fmt.Println("\n------------------------Accesarea elementelor in slice")
	var accesSlice = []int{10, 20, 30, 40}
	fmt.Println(accesSlice[0])
	fmt.Println(accesSlice[1])
	fmt.Println(accesSlice[0:1])

	// stergerea elementului dintr-un slice
	fmt.Println("\n------------------------Stergere element din slice")
	var deleteSlice = []string{"Romania", "Spania", "Italia", "Germania"}
	deleteSlice = RemoveIndex(deleteSlice, 1)

	fmt.Println(deleteSlice)

	// parcugere slice utilizand for loop
	var countrySlice = []string{"Romania", "Spania", "Italia", "Germania"}

	fmt.Println("\n--------------------Exemplu 1 de parcurgere-----------")
	for index, element := range countrySlice {
		fmt.Println(index, "--", element)
	}

	fmt.Println("\n--------------------Exemplu 2 de parcurgere-----------")
	for _, value := range countrySlice {
		fmt.Println(value)
	}

	fmt.Println("\n--------------------Exemplu 3 de parcurgere-----------")
	j := 0
	for range countrySlice {
		fmt.Println(countrySlice[j])
		j++
	}

	// adaugarea unui slice la unul existent
	fmt.Println("\n--------------------Adaugare slice la unul existent")
	var slice1 = []string{"India", "Japonia", "Canada"}
	var slice2 = []string{"Australia", "Rusia"}

	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
