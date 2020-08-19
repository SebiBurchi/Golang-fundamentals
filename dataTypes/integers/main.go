package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//numere: intregi(int), numere cu virgula(float, complex)
	//intregi: pot fi numere intregi semnate sau nesemnate(signed or unsigned)

	//numere intregi semnate pot fi de 5 tipuri: int, int8, int16, int32, int64

	/*int: poate avea dimensiune de 32 de biti sau 64 in functie de platforma pe care ruleaza
	pe 32 de biti range-ul este de la -2^31 pana la 2^31-1
	pe 64 de biti range-ul este de la -2^63 pana la 2^63-1
	*/

	var aInt int // declararea unui int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aInt))
	fmt.Printf("tipul aInt este %s\n", reflect.TypeOf(aInt))

	bInt := 2 //declararea unui int fara specificarea tipului (default-ul va fi tot int)
	fmt.Printf("tipul bInt este %s\n", reflect.TypeOf(bInt))

	/*int8*: poate avea dimensiune de 8 biti, respectiv 1 byte
	range-ul poate fi de la -2^7 pana la 2^7 -1
	*/

	var aInt8 int8 = 2 //declararea unui int8
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aInt8))
	fmt.Printf("tipul aInt8 este %s\n", reflect.TypeOf(aInt8))

	/*int16: poate avea dimensiunea de 16 biti, respectiv 2 bytes
	range-ul poate fi de la 2^15 la 2^15 -1
	*/
	var aInt16 int16 = 2 //declararea unui int16
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aInt16))
	fmt.Printf("tipul aInt16 este %s\n", reflect.TypeOf(aInt16))

	/*int32: poate avea dimensiunea de 32 de biti, respectiv 4 bytes
	range-ul poate fi de la 2^31 la 2^31 -1
	*/
	var aInt32 int32 = 2 //declararea unui int16
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aInt32))
	fmt.Printf("tipul aInt32 este %s\n", reflect.TypeOf(aInt32))

	/*int64: poate avea dimensiunea de 64 de biti, respectiv 4 bytes
	range-ul poate fi de la 2^63 -1 la 2^63
	*/
	var aInt64 int64 = 2 //declararea unui int16
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aInt64))
	fmt.Printf("tipul aInt64 este %s\n", reflect.TypeOf(aInt64))

	//numere intregi nesemnate pot fi de 5 tipuri: uint, uint8, uint16, uint32, uint64
	/*pot avea dimensiuni de 8, 16, 32 respectiv 64 de biti
	  contin doar numere pozitive >=0
	*/

	/*uint: poate avea dimensiunea de 32 de biti sau 64 in functie de platforma pe care ruleaza
	 */
	var aUint uint //declararea unui uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aUint))
	fmt.Printf("tipul aUint este %s\n", reflect.TypeOf(aUint))

	/*uint8: poatea avea dimensiunea de 8 biti/1 byte
	range-ul este de la 0 la 255 sau de la 0 la 2^8 -1
	*/
	var aUint8 uint8 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aUint8))
	fmt.Printf("tipul aUint8 este %s\n", reflect.TypeOf(aUint8))

	/*uint16: poate avea dimensiunea de 16biti/2 bytes
	range-ul este de la 0 la 2^16 -1*/
	var aUint16 uint16 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aUint16))
	fmt.Printf("tipul aUint16 este %s\n", reflect.TypeOf(aUint16))

	/*uint32: poate avea dimensiunea de 32 biti/4 bytes
	range-ul poate fi de la 0 la 2^32 -1*/
	var aUint32 uint32 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aUint32))
	fmt.Printf("tipul aUint32 este %s\n", reflect.TypeOf(aUint32))

	/*uint64: poate avea dimensiunea de 64 de biti/8 bytes
	range-ul poate fi de la 0 la 2^64 -1
	*/
	var aUint64 uint64 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(aUint64))
	fmt.Printf("tipul aUint64 este %s\n", reflect.TypeOf(aUint64))

	/*byte si rune: alias pentru uint8 si int32
	sunt folosite pentru a reprezenta caractere deoarece Golang nu are tipul char
	byte reprezinta caractere ASCII si run caractere Unicode
	in momentul in care declaram o variabila de tip caracter, Go va face type inference catre rune
	*/
	var firstLetter = 'a' // type inference catre rune
	var lastLetter byte = 'z'
	var specCharacter rune = '#'
	fmt.Printf("firstLetter = %d\n", firstLetter)
	fmt.Printf("lastLetter = %d\n", lastLetter)
	fmt.Printf("specCharacter = %U\n", specCharacter)

}
