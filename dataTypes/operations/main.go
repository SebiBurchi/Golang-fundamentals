package main

import "fmt"

func main() {

	//operatii cu numere
	//1. operatii aritmetice
	var a, b = 35, 7
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a mod b = ", a%b)

	//2. operatii de asignare
	var x, y = 15, 25
	x += y // x = x + y
	fmt.Println("x= ", x)

	x = 50
	x -= y // x = x - y
	fmt.Println("x= ", x)

	x = 2
	x *= y // x = x * y
	fmt.Println("x= ", x)

	x = 100
	x /= y // x = x / y
	fmt.Println("x= ", x)

	x = 40
	x %= y // x = x % y
	fmt.Println("x= ", x)

	//3. operatii de incrementare
	var a1, a2 = 2, 5
	a1++ // a1 = a1 + 1
	fmt.Println("a1= ", a1)

	a2++ // a2 = a2 + 1
	fmt.Println("a2= ", a2)

	a2-- // a2 = a2 - 1
	fmt.Println("a2= ", a2)

	//4. operatii de comparare
	var c, d = 15, 25
	fmt.Println(c == d)
	fmt.Println(c != d)
	fmt.Println(c < d)
	fmt.Println(c > d)
	fmt.Println(c >= d)

	//5. operatii logice
	var t1, t2, t3 = 10, 20, 30
	fmt.Println(t1 < t2 && t1 > t3)
	fmt.Println(t1 < t2 || t1 > t3)
	fmt.Println(!(t1 == t2 && t1 > t3))

	//6. operatii la nivel de bit
	var v1 uint = 9  //0000 1001
	var v2 uint = 65 //0100 0001
	var v3 uint
	v3 = v1 & v2
	fmt.Println("v1 & v2 =", v3)

	v3 = v1 | v2
	fmt.Println("v1 | v2 =", v3)

	v3 = v1 << 1
	fmt.Println("v1 << 1 =", v3)

	v3 = v1 >> 1
	fmt.Println("v1 >> 1 =", v3)

}
