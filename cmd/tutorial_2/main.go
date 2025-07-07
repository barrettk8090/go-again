package main

import "fmt"

func main() {
	// [0,0,0]
	var intArr [3]int32
	// Because int32 = 4 bytes of mem, and there are 3 int32s,
	// go initializes array with 12 bytes of contiguous memory
	intArr[0] = 1
	intArr[1] = 123
	intArr[2] = 666
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Printing location in mem of each element
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Shorthand init
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// Or can use [...] which will imply the length of the array based on total elements
	intArr3 := [...]int32{4,5,6}
	fmt.Println(intArr3)
}
