package main

import "fmt"

type gasEngine struct {
	miles_per_gallon uint8 // miles per gallon
	gallons          uint8
	owner            owner
}

// well here is a class method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.miles_per_gallon
}

type electricEngine struct {
	miles_per_kwh uint8 // miles per kwh
	kwh           uint8
}

// well here is a class method
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.miles_per_kwh
}

// well here is just a func
func milesLeft(e gasEngine) uint8 {
	return e.gallons * e.miles_per_gallon
}

func canMakeIt(e gasEngine, miles uint8) {
	// well for now it only works for gasEngine!
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("cannot make it bro")
	}
}

type engine interface {
	// method signature
	// takes no param and returns a uint8
	milesLeft() uint8
}

func canEngineMakeIt(e engine, miles uint8) {
	// well for now it only works for all engines!!!!
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("cannot make it bro")
	}
}

type owner struct {
	name string
}

func main() {
	var myEngine gasEngine = gasEngine{miles_per_gallon: 3, gallons: 6, owner: owner{name: "tom"}}
	fmt.Println(myEngine.gallons)    // 0, default value
	fmt.Println(myEngine.owner.name) // 0, default value
	fmt.Println(myEngine.milesLeft())
	fmt.Println(milesLeft(myEngine))

	canMakeIt(myEngine, 19)
	canEngineMakeIt(myEngine, 19)
	canEngineMakeIt(electricEngine{miles_per_kwh: 3, kwh: 6}, 19)

	var myCar = struct { // not reusable
		brand string
		color string
	}{brand: "RR", color: "black"}

	fmt.Println(myCar.brand, myCar.color)
}
