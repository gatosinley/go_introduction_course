package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int { //A mayuscula para publico, minuscula es privado
	return x + y
}

func RepeatString(total int, value string) {
	for i := 0; i < total; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operation, x float64, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y no puede ser 0")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}
	return 0, errors.New("hubo un error")
}

func Split(v int) (int, int) {
	x := v * 4
	y := v - x
	return x, y
}

func MSum(values ...float64) (float64, error) {
	var sum float64
	if len(values) == 0 {
		return 0, errors.New("no tiene valores")
	}
	sum = values[0]
	for _, v := range values[1:] {
		sum += v
	}
	return sum, nil
}
