package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("número: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("u numero: ", myUintVar)

	var myStringVar string
	myStringVar = "mi primer string"
	fmt.Println(myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("verdad?? ", myBoolVar)
	fmt.Println("mi direccion de memoria es: ", &myStringVar) //& se usa para saber la dirección de memoria

	myIntVar2 := 22 //con esto en vez de ponerle var, hace un casteo automatico con :=
	fmt.Println(myIntVar2)

	const MY_FIRST_CONST = 122 //define el tipo según el valor que le damos, igual se puede declarar antes

	const MY_INT_CONST int = 111122

	var intosho int8 = 127
	fmt.Println(intosho)
	fmt.Printf("valor de int: %d ", intosho)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", intosho, unsafe.Sizeof(intosho), unsafe.Sizeof(intosho)*8)

	intValStr, err := strconv.ParseInt("4", 0, 64)
	fmt.Println(err)
	fmt.Printf("type: %T, value: %d\n", intValStr, intValStr)

}
