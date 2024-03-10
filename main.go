package main

import "fmt"

var x, y int = 20, 100
var msg1 string = "Hello"

func main() {
	z := 30.5
	msg2 := "World"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("z:", z)
	fmt.Println(msg1, msg2)
	fmt.Println(float64(x) + z) //number must be same type
	fmt.Println("Money:",y)
}
