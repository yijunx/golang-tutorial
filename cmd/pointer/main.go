package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	// here we need to use a new here so it assigns a memory location
	// thus later we can safely use *p = 10
	var i int32
	var j int32 = 9

	//   variable  value           memory_location
	//   p         0x400000e0a0    0x4000050020   // here p points to the row where memory location of its value
	//                        0    0x400000e0a0
	//                             0x400000e0a1
	//                             0x400000e0a2
	//                             0x400000e0a3
	//   i                    0    0x400000e0a4
	//                             0x400000e0a5
	//                             0x400000e0a6
	//                             0x400000e0a7
	//   j                    9    0x400000e0a8

	fmt.Println(&p) // 0x4000050020, p's location in memory
	fmt.Println(p)  // 0x400000e0a0, p's value
	fmt.Println(*p) // 0, this is the value of p points to
	fmt.Println(&i) // 0x400000e0a4, i's location
	fmt.Println(i)  // 0, i's value
	fmt.Println(&j) // 0x400000e0a8, j's location

	*p = 10
	// reference the value of the pointer
	p = &i
	*p = 3
	fmt.Println(p)
	fmt.Println(i) // gives 3 because we change the value which p references to

	// this is difference with
	var k int32 = 2
	i = k // this is a copy
	i = 3
	fmt.Println(i) // 3
	fmt.Println(k) // 2

	// say we have a slice
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	fmt.Println(&slice[0]) // gives same memory location!
	fmt.Println(&sliceCopy[0])
	// slices contain pointers to underlying array

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("location of thing1 is %p\n", &thing1) //0x40000aa000

	var result [5]float64 = square(thing1)
	fmt.Printf("result is %v\n", result)
	fmt.Printf("thing1 is %v\n", thing1)

	var result2 [5]float64 = squareUsingPointer(&thing1)
	fmt.Printf("result2 is %v\n", result2)
	fmt.Printf("thing1 is %v\n", thing1) // thing1 changed!!!

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("location of thing2 is %p\n", &thing2) //0x40000aa030
	// the location is different with thing1..
	// thus here we have doubled our memory usage
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func squareUsingPointer(thing3 *[5]float64) [5]float64 {
	// here thing3 is a location which has to points to a array of float64 of size5
	fmt.Printf("location of thing3 is %p\n", thing3) // same location of thing1 !
	// the location is different with thing1..
	// thus here we have doubled our memory usage
	for i := range thing3 {
		thing3[i] = thing3[i] * thing3[i]
	}
	return *thing3 // here we return the dereferenced, not a location
}
