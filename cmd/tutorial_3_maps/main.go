package main

import "fmt"

func main() {
	// key is string, value is uint8
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Printf("Adam is %v years old", myMap2["Adam"])
	// A map will always return something, even if a key does not exist
	fmt.Print("\n", myMap2["Jason"])
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
}
