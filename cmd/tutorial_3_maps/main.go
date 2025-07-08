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
	// delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	intSlice := []int32{4, 5, 6}

	// Iterating over arrays and slices
	for i, v := range intSlice {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// "While" loops in go - bit different!
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	// We can also eliminate the condition

	var j int = 0
	for {
		if j >= 10 {
			break
		}
		fmt.Println(j)
		j = j + 1
	}

	// OR
	for k := 0; k < 10; k++ {
		fmt.Println(k)
	}
}
