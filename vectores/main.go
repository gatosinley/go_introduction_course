package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("bienvenido a los vectores")
	var myArray [5]int
	myArray[1] = 8
	myArray[2] = 5
	myArray[3] = 1
	fmt.Println(myArray)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArray, unsafe.Sizeof(myArray), unsafe.Sizeof(myArray)*8)

	var mySlice []int
	fmt.Printf("size: %d, value: %v\n", len(mySlice), mySlice)

	mySlice = append(mySlice, 10, 20, 30, 40, 50)
	mySlice[4] = 12
	fmt.Printf("size: %d, value: %v\n", len(mySlice), mySlice)

	mySlice = append(mySlice, 10, 20, 30)
	fmt.Printf("size: %d, value: %v\n", len(mySlice), mySlice)

	sliceTest := myArray[2:4] //del indice 2 al 4 pero no incluye el 4
	fmt.Printf("size: %d, value: %v\n", len(sliceTest), sliceTest)

	sliceMake := make([]int, 3)
	fmt.Println(sliceMake)

}
