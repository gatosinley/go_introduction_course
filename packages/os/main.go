package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("myFile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //termina el programa
	}

	data := make([]byte, 200)
	c, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("read %d bytes: %q\n", c, data[:c])

	fmt.Println(os.Getenv("USERNAME"))

	os.Setenv("MI_ENV", "mi valor")

	fmt.Println(os.Getenv("MI_ENV"))

	dir, _ := os.Getwd() //directorio en el que se encuentra
	fmt.Println("dir: ", dir)

	os.Setenv("ENV_EXISTS", "")

	env1 := os.Getenv("ENV_EXISTS")
	env2 := os.Getenv("ENV_NOT_EXISTS")
	fmt.Println(env1)
	fmt.Println(env2)

	env, ok := os.LookupEnv("ENV_EXISTS") //retorna un booleano diciendo si existe o no la variable
	fmt.Println(env, ok)
	env, ok = os.LookupEnv("ENV_NOT_EXISTS")
	fmt.Println(env, ok)

	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "sarasa")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	// ${var} or $var
	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME") //variables que se modifican como en js
	fmt.Println(dbURL)
}
