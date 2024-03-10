package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Go", "Python", "Java"}
	fmt.Println("All the courses:", courseName)
	courseName = append(courseName, "JavaScript", "HTML", "CSS", "C++", "Ruby", "C#") //add element to the end of the array
	fmt.Println("Updated courses:", courseName)

	// Accessing elements in an array using index
	fmt.Println("\nAccessing elements in an array using index")
	fmt.Printf("First Course: %s\n", courseName[0]) // Go
	fmt.Printf("Last Course: %s\n", courseName[5])  // C#

	// Slicing an array to access a subset of the array
	fmt.Println("\nSlice an array to access a subset of the array")
	firstTwo := courseName[:2]    // [0,1]
	lastThree := courseName[6:]   // [6,7,8]
	allButFirst := courseName[1:] // [1,2,3,4,5,6,7,8]
	webCourse := courseName[3:6]  // [3,4,5] start from 3 (is inclusive) and end at 5 (not inclusive)
	fmt.Println("First two courses:", firstTwo)
	fmt.Println("Last three courses:", lastThree)
	fmt.Println("All but first course:", allButFirst)
	fmt.Println("Web courses:", webCourse)

}
