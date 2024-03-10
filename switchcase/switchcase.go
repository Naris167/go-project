package main

import (
	"fmt"
)

func main() {
	// input := 2
	var input string
	fmt.Scanf("%d", &input)

	switch input {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Printf("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("invalid")
	}

}
