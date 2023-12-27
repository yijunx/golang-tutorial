package main // tells the compiler where the first things starts
import (
  "errors"
  "fmt"
)

func main() {
  fmt.Println("hello world!")

  var intNum int = 32767 * 2
  fmt.Println(intNum + 1)

  var floatNum float64 = 123.4
  fmt.Println(floatNum)

  var resultOfAdd float32 = float32(floatNum) + float32(intNum)
  fmt.Println(resultOfAdd)

  var intNume1 int = 3
  var intNume2 int = 2
  fmt.Println(intNume1 / intNume2)

  var myString string = "Hello \nWorld"

  var anotherString string = `hello
world`

  fmt.Println(myString, anotherString)

  fmt.Println(len("adf哈哈")) // gives nine,  // number of utf8 bytes

  var myRune rune = 'a'
  fmt.Println(myRune)

  var myBoolean bool = false
  fmt.Println(myBoolean) // default is faulse

  const myConst string = "constant"
  fmt.Println(myConst)

  printMe(myConst)

  var numerator int = 11
  var denominator int = 0

  var result, remainder, err = intDivision(numerator, denominator)
  if err != nil {
    fmt.Printf((err.Error()))
  } else if remainder == 0 {
    fmt.Printf("The result of the int division is %v", result)
  } else {
    fmt.Printf("The result of the int division is %v with remainder %v", result, remainder)
  }

  switch { // break is auto applied!
  case err != nil:
    fmt.Printf((err.Error()))
  case remainder == 0:
    fmt.Printf("The result of the int division is %v", result)
  default:
    fmt.Printf("The result of the int division is %v with remainder %v", result, remainder)
  }

  switch remainder {
  case 0:
    fmt.Print("And the division is exact!")
  }

  fmt.Println("Here comes array part")
  playWithArray()
}

func printMe(printValue string) {
  fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
  // to handle division by zero
  var err error
  if denominator == 0 {
    err = errors.New("you cant divide by zero")
    return 0, 0, err
  }

  var result int = numerator / denominator
  var remainder int = numerator % denominator
  return result, remainder, err
}

func playWithArray() {
  // []Arrays are
  // Fixed Length
  // Same Type
  // Indexable
  // Contiguous in Memory!

  var intArr [3]int32

  intArr[1] = 123

  fmt.Println(intArr[0])
  fmt.Println(intArr[1:3])

}

