package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	// %v value, %T type
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"s", "t", "r", "i", "n", "g", "y"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
		fmt.Println(strBuilder)
	}
	fmt.Printf("%v, %T", strBuilder, strBuilder)
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
