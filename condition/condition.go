package main

import "fmt"

func main() {
	myMoney := 1000
	if myMoney > 100 {
		println("I'm rich")
	} else if myMoney == 100 {
		println("I'm not rich")
	} else {
		println("I'm poor")
	}

	var score int
	fmt.Print("Enter your score: ")
	fmt.Scanf("%d",&score)

	if score >= 80 {
		fmt.Printf("You score is %d so you got grade A", score)
	} else if score >= 70 {
		fmt.Printf("You score is %d so you got grade B", score)
	} else if score >= 60 {
		fmt.Printf("You score is %d so you got grade C", score)
	} else if score >= 50 {
		fmt.Printf("You score is %d so you got grade D", score)
	} else {
		fmt.Printf("You score is %d so you got grade F", score)
	}
}
