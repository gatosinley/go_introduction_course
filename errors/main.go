package main

import (
	"errors"
	"fmt"

	"github.com/gatosinley/go_introduction_course/errors/operator"
)

func main() {

	var err error
	err = errors.New("my new Error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("my format error String: %s, number: %d", "my String", 23)

	fmt.Println(err2)
	fmt.Println(err2.Error())

	defer func() {
		fmt.Println("In main defer")
		r := recover()

		if r != nil {
			fmt.Println("There arent any results")
			fmt.Println("Recovered in ", r)
			fmt.Printf("error %T\n", r)
		}
	}()

	x := 4
	y := 0
	z := operator.Div(float64(x), float64(y))

	fmt.Println("result")
	fmt.Println("X es: ", z)
}
