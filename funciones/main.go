package main

import (
	"fmt"

	"github.com/gatosinley/go_introduction_course/funciones/function"
)

func main() {
	suma := function.Add(3, 4)

	fmt.Println(suma)

	function.RepeatString(1, "hola")

	valor, err := function.Calc(function.SUM, 3, 6)

	fmt.Println("valor: ", valor, " Error: ", err)

	valor, err = function.Calc(function.SUB, 10, 8)

	fmt.Println("valor: ", valor, " Error: ", err)

	valor, err = function.Calc(function.DIV, 10, 0)

	fmt.Println("valor: ", valor, " Error: ", err.Error())

	x, y := function.Split(20)

	fmt.Println("x: ", x, " y: ", y)

	sum, err := function.MSum(1, 4, 3, 6, 7, 77, 7)
	fmt.Println("total suma: ", sum, " error?: ", err)

	fun := function.FactoryOperation(function.SUM)
	v := fun(2, 3)
	fmt.Println("suma ", v)
}
