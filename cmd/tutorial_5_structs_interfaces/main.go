package main

import "fmt"

type gasEngine struct {
	mpg     uint16
	gallons uint16
	owner
}

type electricEngine struct {
	mpkwh uint16
	kwh   uint16
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint16 {
	return e.gallons * e.mpg
}

func (e gasEngine) getOwner() owner {
	return e.owner
}

func (e electricEngine) milesLeft() uint16 {
	return e.mpkwh * e.kwh
}

func (e electricEngine) getOwner() owner {
	return e.owner
}

type engine interface {
	milesLeft() uint16
	getOwner() owner
}

func canIMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Printf("%v can make it there!!\n", e.getOwner().name)
	} else {
		fmt.Printf("%v needs to fuel up first!\n", e.getOwner().name)
	}
}

func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	var myEngine gasEngine = gasEngine{25, 15, owner{name: "Barrett"}}
	var liamEngine electricEngine = electricEngine{20, 20, owner{name: "Liam"}}
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())
	fmt.Printf("Total miles left in the battery: %v\n", liamEngine.milesLeft())
	canIMakeIt(myEngine, 600)
	canIMakeIt(liamEngine, 399)
}
