package main

import "fmt"

/*
key = product name (string)
value = price (float64)
*/

//when you don't know what the map will contain, you can initialize it with make
var product = make(map[string]float64)

//when you know what the map will contain, you can initialize it with values
var testProduct = map[string]float64{
	"apple":  5.5,
	"banana": 3.4,
}

func main() {
	//another way to initialize a map is by using literal syntax:
	courseName := map[string]string{
		"ITE101": "Intro to Computer",
		"ITE222": "Programming 2",
		"ITE343": "Android Development",
	}

	//show empty map
	fmt.Println("Product:", product)

	//show non-empty map
	fmt.Println("Food:", courseName)

	//add values to the map using map syntax
	product["orange"] = 2.5
	product["grapefruit"] = 3.0
	product["kiwi"] = 4.0

	//accessing a value in the map returns the zero value of its type if no 'value' is present for the 'key'
	fmt.Println("Apple:", testProduct["apple"]) //prints 5.5
	fmt.Println("Grapefruit:", product["grapefruit"]) //prints 3.0
	fmt.Println("Orange:", product["orange"]) //prints 2.5
	fmt.Println("Banana:", testProduct["banana"]) //prints 3.4

	value1 := product["kiwi"]
	fmt.Println("Kiwi:", value1) //prints 4.0

	//delete a key/value pair from the map
	delete(product, "orange")
	fmt.Println("After deleting orange:", product)

	//update a value in the map
	testProduct["apple"] = 10.0
	fmt.Println("Updated apple price:", testProduct["apple"])

	//test for existence of a key in the map
	if _, ok := product["kiwi"]; ok {
		fmt.Println("Kiwi is in the map")
	} else {
		fmt.Println("Kiwi is not in the map")
	}

	//iterate over all keys and values in the map
	for k, v := range product {
		fmt.Println(k, v)
	}


}
