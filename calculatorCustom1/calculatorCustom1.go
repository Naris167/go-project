package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

var OperatorList = map[string]func(float64,float64)float64{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}	


func main() {
	fmt.Println("Welcome to Calculator")
	

	for{
		fmt.Println("Please choose one of the following options:")
		fmt.Println("\n Type '1' to use the applicaiton")
		fmt.Println(" Type '0' to exit the application")
		option := getInput(" Your option: ")

		switch option {
		case 0:
			return
		case 1:
			
			time := getInput("  How many number you want to calculate?\n  Number of times = ")
			num := make([]float64, 0, int(time))
			op := make([]string, 0, int(time)-1)
	
			for i := 0; i < int(time); i++{
				num = append(num, getInput(fmt.Sprintf("   Value %v: ", i+1)))
				if i != int(time)-1{
					op = append(op, getOperator(fmt.Sprintf("   Operator %v: ", i+1)))
				}	
			}
			
			result := num[0]

			for i := 0; i < len(op); i++ {
				result = OperatorList[op[i]](result,num[i+1])
			}

			fmt.Printf("\n  Restul: %v\n\n", result)

		default:
			fmt.Println(" Invalid option")
			continue
		}
	}
}

func getInput(prompt string) float64 {
	for {
		fmt.Printf("%v", prompt)
		input, _ := reader.ReadString('\n')
		number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil{
			fmt.Printf("   `%s` is not a valid number. Please enter the number only\n", strings.TrimSpace(input))
			continue
		}
		return number
	}
}

func getOperator(prompt string) string{
	for {
		fmt.Printf("%v", prompt)
		input, _ := reader.ReadString('\n')
		op := strings.TrimSpace(input)
		if _ , isExists := OperatorList[op]; isExists{
			return op
		} else {
			fmt.Printf("   `%s` is not a valid operator. Please enter the valid operator only", strings.TrimSpace(input))
			continue
		}
	}
}

func add (a, b float64) float64 {
	return a + b
}

func subtract (a, b float64) float64 {
	return a - b
}

func multiply (a, b float64) float64 {
	return a * b
}

func divide (a, b float64) float64 {
	return a / b
}



















