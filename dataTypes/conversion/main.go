package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Conversia tipurilor de date
	*/

	//1. conversia intre tipuri de intregi
	var index int8 = 15
	var bigIndex int32

	bigIndex = int32(index) // conversie dintr-un tip mai mic intr-un tip mai mare
	fmt.Println(bigIndex)

	fmt.Printf("index data type: %T\n", index)
	fmt.Printf("bigIndex data type: %T\n", bigIndex)

	//conversie dintr-un tip mai mare in unul mai mic
	var big int64 = 64
	var little int8

	little = int8(big)
	fmt.Println(little)

	//conceptul de wraparound: un tip de date este convertit intr-un alt tip prea mic si nu poate fi suportat
	var bigWrap int64 = 129
	var littleWrap = int8(bigWrap)
	fmt.Println(littleWrap)

	//2. conversia de la intregi la float
	var xConvert int64 = 57
	var yConvert float64 = float64(xConvert)
	fmt.Printf("%.2f\n", yConvert)

	//3. conversia de la float la intregi
	var fConvert float64 = 390.8
	var iConvert int = int(fConvert)
	fmt.Printf("iConvert = %d\n", iConvert)

	//4. conversia de la numere la string
	aConvert := strconv.Itoa(12)
	fmt.Printf("%q\n", aConvert)

	//5. conversia de la string la intreg
	bConvert, _ := strconv.Atoi("50")
	fmt.Println(bConvert)

	//6. conversia de la string la bytes
	aBytes := "sir de caractere"
	bBytes := []byte(aBytes)
	fmt.Println(bBytes)

}
