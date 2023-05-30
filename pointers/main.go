package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("addrs: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("addrs: %p\n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	i = 22

	fmt.Println(*iP)

	*iP = 11
	fmt.Println(i)

	myVar := 30
	fmt.Printf("addrs: %p, values: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)
	incrementP(&myVar)
	fmt.Println(myVar)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("addrs: %p, values: %v\n", mySlice, mySlice)
	fmt.Printf("addrs 1: %p, addrs 2: %p, addrs 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Printf("addrs: %p, values: %v\n", mySlice, mySlice)

	//myStruct := &MyStruct{1234, "Test"} // se declara como puntero
	myStruct := MyStruct{1234, "Test"}
	fmt.Printf("addrs: %p\n", &myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(&myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	myStruct.Set("test method")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	myStruct.SetP("test method 2")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(v []int) {
	fmt.Printf("addrs: %p\n", v)
	fmt.Printf("addrs 1: %p, addrs 2: %p, addrs 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
	v = append(v, 9, 8) //estos no se verán reflejados, mueren en el scope local
}

func updateStruct(v *MyStruct) {
	fmt.Printf("addrs: %p\n", v)
	v.ID = 999
	v.Name = "updated"
}
