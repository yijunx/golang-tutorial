package main

import (
	"fmt"
	"time"
)

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)

	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	myMap2["Olivia"] = 23

	fmt.Println(myMap2)

	fmt.Println(myMap2["Tom"]) // here you get the default value of that type, no key error

	var age, ok = myMap2["Tom"]

	if ok {
		fmt.Printf("Toms age is %v", age)

	} else {
		fmt.Println("Invalid name")
	}

	delete(myMap2, "Name not there") // here you can delete anyhow, nice!

	for name := range myMap2 {
		fmt.Printf("Name is %v and age is %v \n", name, myMap2[name])
	}

	for name, age := range myMap2 {
		fmt.Printf("Name is %v and age is %v \n", name, age)
	}

	var myArr [3]int32 = [...]int32{1, 2, 3}

	for i, v := range myArr {
		fmt.Printf("Index is %v and value is %v \n", i, v)
	}

	// while loop
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	// init / condition / post (i--, i++, i += 10, i -= 10, i *= 10, i /= 10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var testSlice2 = make([]int, 0, 1000)
	fmt.Printf("len is %d, cap is %d \n", len(testSlice2), cap(testSlice2))

	checkTimeLoop()

}

func checkTimeLoop() {
	var n int = 1_000_000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
