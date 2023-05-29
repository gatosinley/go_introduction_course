package main

import "fmt"

func main() {

	mapa := make(map[int]string)
	mapa[1] = "A"
	mapa[2] = "B"
	mapa[3] = "C"
	fmt.Println(mapa)
	fmt.Println(mapa[1])

	delete(mapa, 2)
	fmt.Println(mapa)

	mapa[1] = "A2"
	fmt.Println(mapa)
	v, ok := mapa[3]
	fmt.Println("The value: ", v, " existe?: ", ok)
	fmt.Println(mapa[19])
	_, ok = mapa[19] //_ es para decir que no usaremos ese valor
	fmt.Println(" existe?: ", ok)

	//sin make
	mapa2 := map[int]string{
		4: "A",
		5: "F",
		6: "L",
	}
	fmt.Println(mapa2)
}
