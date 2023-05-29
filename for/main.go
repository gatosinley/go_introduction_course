package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	sum = 1

	for sum < 100 {
		sum++
	}

	for {
		if sum > 1000 {
			break
		}
		sum++
	}

	array := []int{1, 2, 3, 4, 5, 6}
	for i := range array {
		fmt.Println("index: ", i, " - value", array[i])
	}
	fmt.Println(sum)

	mapa := map[string]float64{
		"A": 12.3,
		"B": 22.4,
		"C": 10.8,
	}

	for key, value := range mapa {
		fmt.Println("key: ", key, " Value: ", value)
	}

	mapaArray := map[string][]int{
		"A": nil,
		"B": {1, 2, 5, 4, 8, 40},
		"Z": {9, 5, 2, 54, 6, 3},
	}

	for key, value := range mapaArray {
		fmt.Println("key: ", key)
		for _, v := range value {
			fmt.Println("value ", v)
		}
	}

}
