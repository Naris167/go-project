package main

import (
	"fmt"
)

const count = 15

func main() {
	// For Loop Example
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("i =", i)
	// 	i++
	// }

	// for j := 0; j < count; j++ {
	// 	fmt.Println("j =", j)
	// }

	// While Loop Example
	var input string
	for {
		fmt.Scanf("%s", &input)
		if input != "exit" {
			fmt.Println("Input =", input)
		} else {
			break
		}
	}

}
