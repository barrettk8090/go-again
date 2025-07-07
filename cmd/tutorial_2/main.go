package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println("intArr:", intArr)

	// Appending + checking the capacity of the slice
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length of the slice is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe NEW length of the slice is %v with NEW capacity %v\n", len(intSlice), cap(intSlice))
	// ^^ Capacity of the slice becomes 6

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println("Length: ", len(intSlice3))
	fmt.Println("Capacity: ", cap(intSlice3))
}
