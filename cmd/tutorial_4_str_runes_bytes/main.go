package main

import "fmt"

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	// %v value, %T type
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	
}
