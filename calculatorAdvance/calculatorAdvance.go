package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


// var OperatorList []string

var OperatorList = map[string]func(float64, float64) float64{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func main() {
	// OperatorList = append(OperatorList, "+", "-", "*", "/")
	fmt.Println(" Welcome to Calculator")

	for {
		var input string
		var num1, num2, result float64

		fmt.Println("\n Type '1' to use the applicaiton")
		fmt.Println(" Type '0' to exit the application")
		fmt.Print(" Your option: ")
		/*
		This line is commented out to avoid the error of "unexpected newline in argument list".
		The error is caused by the fact that the fmt.Scan function does not handle
		the newline character properly. It is better to use the bufio package to read input from the user.

		1. `fmt.Scan(&input)` reads input from the standard input buffer until (not include) it encounters whitespace.
		2. It then leaves any subsequent characters in the buffer, including the newline character
		   if the user pressed Enter immediately after their input.
		3. `bufio.NewReader(os.Stdin).ReadString('\n')` reads data from the buffer up to
		   and including the specified delimiter, which in this case is the newline character (\n).
		4. If the first character in the buffer is a newline character (left over from a previous
		   input operation like `fmt.Scan`), `ReadString` will read just that newline character,
		   resulting in an empty string as the input (since the delimiter is removed from the returned string).
		5. The empty string is then passed to `strconv.ParseFloat`, which fails to parse it as a float,
		   leading to an error.
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

			times := getInput("  How many number you want to calculate?\n  Number of times = ")
			nums := make([]float64, 0)
			operators := make([]string, 0)

			/*
			For testing purposes, you can uncomment the following lines to use predefined values.
			Result should be 1000 / 2 + 100 - 20 + 100 = 680
			
			nums = append(nums, 1000, 2, 100, 20, 100)
			operators = append(operators, "/", "+", "-", "+")
			*/

			for i := 0; i < int(times); i++ {
				nums = append(nums, getInput(fmt.Sprintf("   Value %v = ", i+1)))
				if i == int(times)-1{
					break
				}
				operators = append(operators, getOperator(fmt.Sprintf("   Operator %v = ", i+1)))
			}

			// fmt.Println("Get input success")

			result = nums[0] // Start with the first number
			for i := 0; i < int(times)-1; i++ {
				num1 = result
				num2 = nums[i+1]
				op := operators[i]

				opFunc := OperatorList[op]
				result = opFunc(num1, num2)



				// For testing purposes, you can uncomment the following line to see the loop number and result.
				// fmt.Printf("Loop number %v, result = %v\n", i, result)
				
				// switch operator {
				// case "+":
				// 	result = add(num1, num2)
				// case "-":
				// 	result = subtract(num1, num2)
				// case "*":
				// 	result = multiply(num1, num2)
				// case "/":
				// 	result = divide(num1, num2)
				// default:
				// 	panic("  Invalid Operator")
				// }

			}
		default:
			fmt.Println(" Invalid Option")
		}
		fmt.Printf("\n  Restul: %v\n", result)
	}
}

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	for {
		fmt.Printf("%v", prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			/*

			message, _ := fmt.Scanf("%v is not a valid number", prompt)
			panic(message)

			These lines are commented out because:
			1. Misunderstanding with fmt.Scanf: Attempting to use fmt.Scanf to create an error message
			is incorrect as it is for reading and scanning input, not for formatting error messages.
			This misunderstanding led to confusion when you expected it to format an error message.

			2. Immediate Error without Custom Message: Because the conversion fails, the program
			enters the error handling block (if err != nil { ... }). But instead of showing
			a custom error message ("is not a valid number"), it directly goes to a panic because
			the attempt to create a formatted message was incorrect. You saw an immediate error (panic)
			because the code was not correctly generating or displaying the intended custom error message.
			*/
			
			
			// Not use this anymore as it will cause the program to panic immediately if the user enters something other than a number.
			// panic(fmt.Sprintf("'%s' is not a valid number.", strings.TrimSpace(input)))

			fmt.Printf("'%s' is not a valid number. Please enter a valid number.\n", input)
			continue
		}
		return value
	}
}


func getOperator(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		op := strings.TrimSpace(input)

		// for i := 0; i < len(OperatorList); i++ {
		// 	if op != OperatorList[i] {
		// 		continue
		// 	} else {
		// 		return op
		// 	}
		// }

		if _ , exists := OperatorList[op]; exists {
			return op
		} else {
			fmt.Println("  Invalid Operator")
			continue
		}
	}
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
