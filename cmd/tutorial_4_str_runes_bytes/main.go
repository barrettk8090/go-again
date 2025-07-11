package main

import "fmt"

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	// %v value, %T type
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	myRune := 'a'
	myNewString := "a"
	fmt.Printf("\n%v, %T", myRune, myRune)
	fmt.Printf("\n%v, %T", myNewString, myNewString)

}
