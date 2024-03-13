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
	"+":add,
	"-":subtract,
	"*":multiply,
	"/":divide,
}

func main() {
	for{
		fmt.Println("select the option")
		fmt.Println("1 to use application")
		fmt.Println("0 to exit")
		option := getInput("You option: ")

		switch option {
		case 0:
			return
		case 1: 
			time := getInput("how many number you want to calculate\nnumber: ")
			num := make([]float64, 0, int(time))
			op := make([]string, 0, int(time)-1)
		

			for i := 0; i < int(time); i++{
				num = append(num, getInput(fmt.Sprintf("value %v: ", i+1)))
				if i != int(time)-1 {
					op = append(op, getOperator(fmt.Sprintf("opertor %v: ", i+1)))
				}
			}

			result := num[0]

			for i := 0; i < int(time)-1; i++ {
				result = OperatorList[op[i]](result,num[i+1])
			}

			fmt.Printf("Result: %v\n", result)

		default:
			fmt.Println("invalid option")
		}


	}
}

func getInput(prompt string)float64{
	for {
		fmt.Printf("%v",prompt)
		input, _ := reader.ReadString('\n')
		num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil{
			fmt.Println("This is not the correct number")
			continue
		}
		return num
	}
}

func getOperator(prompt string)string{
	for {
		fmt.Printf("%v", prompt)
		input, _ := reader.ReadString('\n')
		op := strings.TrimSpace(input)
		if _ , isExist := OperatorList[op]; isExist{
			return op
		}
		fmt.Println("Invalid operator")
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
