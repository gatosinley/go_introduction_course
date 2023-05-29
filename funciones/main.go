package main

import (
	"fmt"

	"github.com/gatosinley/go_introduction_course/funciones/function"
)

func main() {
	suma := function.Add(3, 4)

	fmt.Println(suma)

	function.RepeatString(5, "hola")
}
