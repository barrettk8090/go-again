package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	// Get the remainder
	fmt.Println(intNum1 % intNum2)

	var myString string = `Hello 
World`
	fmt.Println(myString)

	// Getting the number of bytes in a string
	fmt.Println(len("test")) // 4
	fmt.Println(len("A"))    // 1
	fmt.Println(len("§"))    // 2

	// Getting the actual length of a string using utf9
	fmt.Println(utf8.RuneCountInString("Hello∞")) // 6

	// Runes are weird!
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// Can omit types if you set the value right away. Type inferred
	var myVar = "text"
	fmt.Println(myVar)

	// Remove the var keyword, use shorthand :=
	newVar := "text"
	fmt.Println(newVar)

	// Init multiple vars at once
	var1, var2 := 1, 2
	fmt.Println(var1, var2)
	var3, var4 := 3, "Four"
	fmt.Println(var3, var4)

	// Constants
	const myConst string = "const value"
	fmt.Println(myConst)

	printValue := "hellooo world"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var division, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Print(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", division, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
