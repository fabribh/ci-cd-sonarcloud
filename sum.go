package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
	fmt.Println(times(2, 2))
	fmt.Println(sub(10, 2))
	fmt.Println(div(8, 2))
}

func sum(a int, b int) int {
	return a + b
}

func times(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	return a / b
}
