package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//numere cu virgula
	//1. float32 si float64

	/*float32: poate avea dimensiunea de 32 biti / 4 bytes
	este setul tuturor numerelor cu virgula IEEE-754 pe 32 de biti
	poate fi mai rapid decat float64
	valoarea default este 0.0
	*/
	var aFloat32 float32 = 2.0
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aFloat32))
	fmt.Printf("tipul aFloat32 este %s\n", reflect.TypeOf(aFloat32))

	/*float64: poate avea dimensiunea de 64 biti / 8 bytes
	este setul tuturor numerelor cu virgula IEEE-754 pe 64 de biti
	poate reprezenta numere cu o acuratete mai mare decat float32
	valoarea default este 0.0
	se foloseste cand precizia necesare este mai mare
	*/
	var aFloat64 float64 = 3.0
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aFloat64))
	fmt.Printf("tipul aFloat64 este %s\n", reflect.TypeOf(aFloat64))

	//default-ul pentru float este float64
	bFloat := 2.3
	fmt.Printf("tipul bFloat este %s\n", reflect.TypeOf(bFloat))

	//2. complex64 si complex128: numere reale cu parte imaginara si parte reala

	/* complex64: atat partea imaginara cat si cea reala pot avea 64 de biti / 8 bytes
	reprezinta setul tuturor numerelor cu partea reala si cea imaginara de tipul float32
	se poate initializa in mai multe feluri
	*/

	//initializare 1 numar complex64
	var aRealFloat32 float32 = 3
	var bImaginarFloat32 float32 = 5
	cComplex64 := complex(aRealFloat32, bImaginarFloat32)

	//initializare 2 numar complex64
	var dComplex64 complex64
	dComplex64 = 4 + 5i

	fmt.Printf("dimensiunea cComplex64 este %d bytes\n", unsafe.Sizeof(cComplex64))
	fmt.Printf("dimensiunea dComplex64 este %d bytes\n", unsafe.Sizeof(dComplex64))
	fmt.Printf("tipul cComplex64 %s\n", reflect.TypeOf(cComplex64))
	fmt.Printf("tipul dComplex64 %s\n", reflect.TypeOf(dComplex64))

	/* complex128: atat partea reala cat si cea imaginara pot avea 128 de biti / 16 bytes
	reprezinta setul tutoror numerelor cu partea reala si cea imaginara de tipul float64
	*/

	//initializare 1 numar complex128
	var aRealFloat64 float64 = 3
	var bRealFloat64 float64 = 5
	cComplex128 := complex(aRealFloat64, bRealFloat64)

	//initializare 2 numar complex128
	var dComplex128 complex128
	dComplex128 = 4 + 5i

	fmt.Printf("dimensiunea cComplex128 este %d bytes\n", unsafe.Sizeof(cComplex128))
	fmt.Printf("dimensiunea dComplex128 este %d bytes\n", unsafe.Sizeof(dComplex128))
	fmt.Printf("tipul cComplex128 %s\n", reflect.TypeOf(cComplex128))
	fmt.Printf("tipul dComplex128 %s\n", reflect.TypeOf(dComplex128))

	//operatii cu numere complexe
	fmt.Println(cComplex128+dComplex128, cComplex128-dComplex128, cComplex128*dComplex128, cComplex128/dComplex128)

}
