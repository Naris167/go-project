package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

var OperatorList = map[string]func(float64, float64) float64{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func main() {
	for {
		fmt.Println("welcome to calculator")
		fmt.Println("type 1 to enter application")
		fmt.Println("type 0 to exit application")
		option := getInput("your option: ")

		switch option {
		case 0:
			return
		case 1:
			time := getInput("How many number you want to calculate?\nNumber of time: ")
			number := make([]float64, 0, int(time))
			op := make([]string, 0, int(time)-1)

			for i := 0; i < int(time); i++ {
				number = append(number, getInput(fmt.Sprintf("value %v: ", i+1)))
				if i != int(time){
					op = append(op, getOperator(fmt.Sprintf("operator %v: ", i+1)))
				}
			}

			result := number[0]

			for i := 0; i < int(time)-1; i++ {
				result = OperatorList[op[i]](result, number[i+1])
			}

			fmt.Printf("Restul is %v", result)

		default:
			fmt.Println("invalid option")
		}

	}
}

func getInput(prompt string) float64 {
	for {
		fmt.Printf("%v", prompt)
		input, _ := reader.ReadString('\n')
		num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("invalid number")
			continue
		}
		return num
	}
}

func getOperator(prompt string) string {
	for {
		fmt.Printf("%v", prompt)
		input, _ := reader.ReadString('\n')
		op := strings.TrimSpace(input)
		if _, isExist := OperatorList[op]; isExist {
			return op
		}
		fmt.Println("invalid operator")
		continue
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
