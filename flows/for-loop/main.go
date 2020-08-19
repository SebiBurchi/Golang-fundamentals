package main

import "fmt"

func main() {

	/*
		for loop este folosit pentru a executa o bucata de cod in mode repetat
		este singura structura de control repetitiva din Golang
		cuprinde 3 parti: initializare, conditie si incremntare

	*/

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	} // Golang nu contine paranteze rotunde pentru for, iar {} sunt obligatorii

	// omiterea partii de initializare
	fmt.Println("\nOmiterea initializarii pentru for loop")
	i := 3
	for ; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}

	// omiterea partii de incrementare
	fmt.Println("\nOmiterea incrementarii pentru for loop")
	i = 2
	for i <= 20 {
		fmt.Printf("%d ", i)
		i *= 2 // i = i * 2
	}

	// omiterea tuturor conditiilor - va duce la un loop infinit
	// for {

	// }

	// break este folosit pentru a termina un loop inainte de termen
	fmt.Println("\nExemplu break")
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("Primul numar pozitiv divizibil cu 3 si 5 este %d\n", num)
			break
		}
	}

	// continue este folosit pentru a stopa iteratia curenta din loop, dupa care se trece la iteratia urmatoare
	fmt.Println("\nExemplu continue")
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
}
