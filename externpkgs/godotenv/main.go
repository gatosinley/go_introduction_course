package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	println("intento 1")
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	if err := godotenv.Load("otherFolder/.var"); err != nil { //load funciona solo la primer vez
		fmt.Println(err)
	}

	println("intento 2")
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	myEnv, err := godotenv.Read("otherFolder/.var")
	fmt.Println(err)
	fmt.Println(myEnv)

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	println("intento 3")
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

}
