package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	var myEngine gasEngine = gasEngine{25, 15, owner{"Barrett"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
}
