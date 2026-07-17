package main

import "fmt"

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld() {
	RunHelloWorld()
}

func Perkalian(a int, b int) int {
	result := a * b
	return result
}
