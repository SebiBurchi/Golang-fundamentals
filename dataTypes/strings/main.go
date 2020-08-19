package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	/*
		Go suporta doua modalitati de declrare a string-urile: double-quote si back-quote
		Valoarea zero pentru string-uri este ""	``
		Pot fi concatenate utilizat operatori de genul + sau +=
		Pot fi comparate utilizand operatori de genul == != > < >= <=

	*/

	myStringDouble := "Primul string"           //crearee si initializare string utilizand metoda double-quote
	myStringEscape := "Primul string!\nTelacad" //escaparea funtioneaza doar pentru metoda double-quote

	myStringBack := `Al doilea string`                 //create si initializare string utilizand metoda back-quote
	myStringBackEscape := `Al doilea string!\nTelacad` //escaparea nu functioneaza pentru metoda back-quote

	fmt.Println("myStringDouble: ", myStringDouble)
	fmt.Println("myStringEscape: ", myStringEscape)
	fmt.Println("myStringBack: ", myStringBack)
	fmt.Println("myStringBackEscape: ", myStringBackEscape)

	//string este imutabil: read-only (incercam sa ii schimbam valoarea -> eroare de compilare)
	//myStringDouble[1] = 'A'

	//lungimea unui string utilizand functia len()
	length1 := len(myStringDouble)
	fmt.Println("Length 1:", length1)

	//lungimea unui string utilizand functia RuneCountInString()
	length2 := utf8.RuneCountInString(myStringDouble)
	fmt.Println("Length 2:", length2)

	//concatenare string-uri
	var concatenareString = myStringDouble + " " + myStringEscape
	concatenareString += "!" //concatenareString = concatenareString + "!"
	fmt.Println("Concatenare string: ", concatenareString)

	//comparare string-uri
	fmt.Println(myStringDouble == "Primul string")
	fmt.Println(myStringDouble > myStringEscape)
	fmt.Println(myStringDouble != myStringEscape)

	//substring
	var mySubString = myStringDouble[1:5]
	fmt.Println("Substring: ", mySubString)
	fmt.Println("Substring[0]", mySubString[0]) //va afisa codul ASCII sau Unicode pentru primul caracter

	//cateva functii specifice string
	fmt.Println("Contains: ", strings.Contains(myStringDouble, "Primul"))
	fmt.Println("HasPrefix: ", strings.HasPrefix(myStringDouble, "Pri"))
	fmt.Println("Index: ", strings.Index(myStringDouble, "r"))

}
