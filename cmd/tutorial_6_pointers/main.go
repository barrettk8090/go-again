package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 // 0
	fmt.Printf("The value p points to is %v, in memory location: %v", *p, p)
	fmt.Printf("\nThe value of i is: %v, in memory location: %v", i, &i)

	p = &i
	fmt.Printf("\nThe value p points to is still %v, but now in memory location: %v", *p, p)

	*p = 1
	fmt.Printf("\nBecause the value of p is now %v, the value of i is also %v", *p, i)
}
