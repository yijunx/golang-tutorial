package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	// As for the encoding/json package though -
	// it cannot see them if not exported.
	// You need to make all of your fields visible to the encoding/json package
	// by making them all start with an uppercase letter, thereby exporting them

	Name   string // well it has to be capital N, so it gets exported
	Email  string // exported
	gender string // Not exported
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error loading file %v\n", filePath)
	}
	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}

type gasEngine struct {
	gallons          float32
	miles_per_gallon float32
}

type electricEngine struct {
	kwh           float32
	miles_per_kwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	// var intSlice = []int{1, 2, 3}
	// // additional param in the square bracket!
	// fmt.Println(sumSlice[int](intSlice))

	// var intSlice2 = []int{}
	// fmt.Println(isEmpty(intSlice2))

	var x []contactInfo = loadJSON[contactInfo]("cmd/generics/contactInfo.json")
	var y []purchaseInfo = loadJSON[purchaseInfo]("cmd/generics/purchaseInfo.json")

	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", y)

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons:          12.4,
			miles_per_gallon: 23,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:           34,
			miles_per_kwh: 123,
		},
	}

	fmt.Printf("%+v and %+v \n", gasCar, electricCar)

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	// check if slice of any time is empty!!
	return len(slice) == 0
}
