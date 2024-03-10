package main

import "fmt"

//array need to specify the size at compile time, slice does not.

var pName [4]string //don't need to put value now

func main() {
	price := [4]float32{15.5, 25.5, 35.5, 45.5}//must put value now

	pName[0] = "Shirt"
	pName[1] = "Pant"
	pName[2] = "Shoe"
	pName[3] = "Hat"

	fmt.Println(pName)
	fmt.Println(price)
}
