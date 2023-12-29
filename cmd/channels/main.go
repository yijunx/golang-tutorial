// channel, go routines that passes around information
// channels hold data, thread safe, avoid data races, listen for data, block code execution

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// var c = make(chan int, 3) // this channel holds a single int
	// var bufferChannel = make(chan int, 5)
	// c <- 1                 // added value 1 to the channel
	// // we are waiting here forever
	// var i = <-c
	// fmt.Println(i)
	// fatal error: all goroutines are asleep - deadlock!
	// thus we need goroutine, they go hand in hand with channel

	// go process(c)
	// code move to print first, then it sees it needs to
	// take out something from c
	// then the process finish, c listens for data
	// then it prints

	// for i := range c {
	// 	// printStuff(<-c)
	// 	time.Sleep(time.Second * 1)
	// 	printStuff(i)

	// }

	// a more realist example
	var chickenChannel = make(chan string) // holds the website we find the sale on
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		// here we spawn 3 go routine, we try to find until
		// a cheap chicken price is found
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	// the send message function is waiting here for value to be added to chanel
	sendMessage(chickenChannel, tofuChannel)

}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Printf("%s sells chicken at %v...\n", website, chickenPrice)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Printf("%s sells tofu at %v...\n", website, chickenPrice)
		if chickenPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel: // this means we found a cheap chicken..
		fmt.Printf("Found a deal on chicken at %s\n", website)
	case website := <-tofuChannel: // this means we found a cheap tofu..
		fmt.Printf("Found a deal on tofu at %s\n", website)
	}

}

func process(c chan int) {
	// c <- 123 // add value 123 to the channel
	defer close(c) // do this stuff before function exists
	for i := 0; i < 5; i++ {
		fmt.Printf("adding number %d to c \n", i)
		c <- i
	}
	fmt.Println("process ends!")
	// close(c) // tells all listener that channel closed, so wont waiting forever -> deadlock
}

func printStuff(stuff int) {
	fmt.Println("Print is triggered")
	fmt.Println(stuff)
}

// result:

// adding number 0 to c
// adding number 1 to c
// Print is triggered
// 0
// Print is triggered
// 1
// adding number 2 to c
// adding number 3 to c
// Print is triggered
// 2
// Print is triggered
// 3
// adding number 4 to c
// Print is triggered
// 4
