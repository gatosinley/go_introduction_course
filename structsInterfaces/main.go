package main

import (
	"fmt"

	"github.com/gatosinley/go_introduction_course/structsInterfaces/interfaces"
	"github.com/gatosinley/go_introduction_course/structsInterfaces/structs"
)

func main() {
	fmt.Println("hola go")

	var p1 structs.Product
	p1.ID = 11
	p1.Name = "Clara"
	fmt.Println(p1)
	/*
		p2 := structs.Product{}
		p2.ID = 123
		p2.Name = "juan"
		fmt.Println(p2)

		p3 := structs.Product{2, "Alberta", structs.Type{1, "A", "aaaaaaa"}, 12555, 15, nil} // el warning es por no poner los key - se deben llenar todos sin key
		fmt.Println(p3)
		typeS := structs.Type{ID: 1, Code: "A", Description: "aaaaaaa"}
		p4 := structs.Product{ID: 5, Name: "Marcos", Type: typeS, Price: 9999, Count: 25, Tags: nil} // con los key puedo no llenar todos los campos
		fmt.Println(p4)

		v, err := json.Marshal(p4)
		fmt.Println("valor p4 ", string(v), " error: ", err) // se agrega string pq el json lo retorna solo como bytes

		fmt.Println("precio total: ", p4.TotalPrice())
		p4.SetName("Almendra")
		p4.AddTags("tag1", "tag2", "tag3", "tag4", "tag5")
		fmt.Println(p4)
	*/
	fmt.Println("Vehiculos")
	carV := interfaces.Car{Time: 160}
	fmt.Println(carV.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "CAR", "MOTORCYCLE", "TRUCK", "asdf"}

	d := 400
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s\n", v)
		veh, err := interfaces.New(v, d)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Printf("distancia: %.2f\n", veh.Distance())
	}

}
