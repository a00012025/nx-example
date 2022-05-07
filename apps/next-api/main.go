package main

import (
	"fmt"
	"os"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("next-api-2"))
	dat, _ := os.ReadFile("libs/ascii/assets/cow.txt")
	fmt.Print(string(dat) + "\n")
	fmt.Println("Hello World!")
}
