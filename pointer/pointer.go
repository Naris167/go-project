package main

import "fmt"

func zeroValue(ivalue int) {
	ivalue = 0
}

func zeroPointer(ipointer *int) {
	*ipointer = 0 //Go to the address of 'original' i and change it to 0
}

func main() {
	i := 1
	fmt.Println("Initial value of i:", i)

	//This will not change the value of i
	zeroValue(i)
	fmt.Println("i from zeroValue:", i)

	//This will change the value of i to 0
	zeroPointer(&i) //Getting the address of i and passing that address to zeroPointer
	fmt.Println("i value from zeroPointer:", i)
	fmt.Println("i address from zeroPointer:", &i)
}
