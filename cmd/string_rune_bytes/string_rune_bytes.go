package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "resume雷猴" // it has non ascii chars!
	fmt.Println(myString)

	var indexed = myString[0]
	fmt.Printf("%v, %T \n", indexed, indexed)
	// we are indexing the underlying byte array!

	for i, v := range myString {
		fmt.Printf("index is %d, value is %v, type is %T\n", i, v, v)
		// 0,1,2,3,4,5,6,9
		// utf-8 uses varied length of bytes for encoding
		// a -> 01100001 (1)
		// 家 -> 11100101 10101110 10110110 (3)
	}

	// now lets try using runes
	var myRunes = []rune("8+9等于17")
	fmt.Printf("the length of myRunes is %v\n", len(myRunes))
	for i, v := range myRunes {
		fmt.Printf("index is %d, value is %v, type is %T\n", i, v, v)
		// 0,1,2,3,4,5,6,9
		// utf-8 uses varied length of bytes for encoding
		// a -> 01100001 (1)
		// 家 -> 11100101 10101110 10110110 (3)
		// rune is just alias for int32
	}

	var strSlice = []string{"s", "u", "hhh"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i] // here it create new string every time, it is inefficient
	}
	fmt.Printf("%v\n", catStr)

	var stringBuilder strings.Builder

	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("%v\n", stringBuilder.String())

}
