package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" Welcome to Calculator")

	for {
		var input string
		var result float64

		fmt.Println("\n Type '1' to use the applicaiton")
		fmt.Println(" Type '0' to exit the application")
		fmt.Print(" Your option: ")
		/*
		This line is commented out to avoid the error of
		"unexpected newline in argument list". The error is
		caused by the fact that the fmt.Scan function does not
		handle the newline character properly.
		It is better to use the bufio package to
		read input from the user.
		*/
		// fmt.Scan(&input) 

		input, err := reader.ReadString('\n')
        if err != nil {
            panic(" Failed to read option")
        }
        input = strings.TrimSpace(input)

		switch input {
		case "0":
			return
		case "1":

			num1 := getInput("\n  Value 1 = ")
			num2 := getInput("  Value 2 = ")

			switch operator := getOperator(); operator {
			case "+":
				result = add(num1, num2)
			case "-":
				result = subtract(num1, num2)
			case "*":
				result = multiply(num1, num2)
			case "/":
				result = divide(num1, num2)
			default:
				panic("  Invalid Operator")
			}
		default:
			fmt.Println(" Invalid Option")
		}

		fmt.Printf("\n  Restul: %v\n", result)
	}
}

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		/*
		This line is commented out to because:
		1. Misunderstanding with fmt.Scanf: Attempting to use fmt.Scanf to create an error message
		is incorrect as it is for reading and scanning input, not for formatting error messages.
		This misunderstanding led to confusion when you expected it to format an error message.

		2. Immediate Error without Custom Message: Because the conversion fails, the program
		enters the error handling block (if err != nil { ... }). But instead of showing
		a custom error message ("is not a valid number"), it directly goes to a panic because
		the attempt to create a formatted message was incorrect. You saw an immediate error (panic)
		because the code was not correctly generating or displaying the intended custom error message.
		*/
		// message, _ := fmt.Scanf("%v is not a valid number", prompt)
		// panic(message)
		
		panic(fmt.Sprintf("'%s' is not a valid number.", strings.TrimSpace(input)))
	}

	return value
}


func getOperator() string {
	fmt.Print("  Choose operator (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
