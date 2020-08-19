package main

import "fmt"

var g = "global"

var test1 = 10

func printLocal() {
	l := "local"
	fmt.Println(l)
}

func printLocalGlobal() {
	l := "local2"
	fmt.Println(l)
	fmt.Println(g)
}

func modifyGlobal() {
	g = "global2"
	fmt.Println(g)
}

func useSameVariableName() {
	test1 := 20
	test2 := 25

	fmt.Println(test1)
	fmt.Println(test2)
}

const (
	testConst  = 20
	testConst2 = "John"
)

func printLocalConst() {
	const testConst3 = 35.5
	fmt.Println(testConst3)
}

func useSameConstName() {
	const testConst = 36
	fmt.Println(testConst)
}

/*func modifyConstValue() {
	testConst = 22
}
	eroare de compilare pentru ca este o constanta
*/

func main() {

	/*
		scopul variabilelor si constantelor - local sau global
		variabilele si constantele locale exista doar in interiorul functiilor unde sunt declarate
		variabilele si constantele globale exista si in afara functiilor in care sunt declarate
	*/

	//1. variabile
	printLocal()
	fmt.Println(g)

	//fmt.Println(l) //eroare de compilare deoarece l este o variabila locala

	printLocalGlobal()

	//modificare variabila globala
	fmt.Println(">>>>>>>>>>>> modificare variabila globala")
	modifyGlobal()
	fmt.Println(g)

	fmt.Println(">>>>>>>>>>>> utilizare acelasi nume pentru variabile locale si globale")
	// utilizam acelasi nume pentru variabile locale si globale
	useSameVariableName()
	fmt.Println(test1)

	//2. constante
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>> constante")
	fmt.Println(testConst)
	fmt.Println(testConst2)

	printLocalConst()

	// utilizare acelasi nume pentru constante locale si globale
	useSameConstName()
}
