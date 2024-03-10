package main

import "fmt"

//Ex: Function that only return value
func hello() {
	fmt.Println("Hello, World!")
}

//Ex: Function that have the parameter and return value
func plus(a int, b int) int {
	return a + b
}

//Ex: Another way to declare the parameter and return value
func plus3Value(a, b, c int) int {
	return a + b + c

}

func main() {
	hello()
	result := plus(1, 2)
	fmt.Println(result)

	result2 := plus3Value(1, 2, 3)
	fmt.Println(result2)

}
