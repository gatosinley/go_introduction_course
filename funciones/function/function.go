package function

import "fmt"

func Add(x int, y int) int { //A mayuscula para publico, minuscula es privado
	return x + y
}

func RepeatString(total int, value string) {
	for i := 0; i < total; i++ {
		fmt.Println(value)
	}
}
