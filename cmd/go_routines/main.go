package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.Mutex{}      // mutual exclusion
var m_rw = sync.RWMutex{} // read write mutual exclusion
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // increase the counter
		go dbCall(i)
	}
	wg.Wait() // we need to wait
	fmt.Printf("total execution time: %v\n", time.Since(t0))
	fmt.Printf("finally, the results are %v\n", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000 // up to 2 seconds per call
	delay = 2000                              // now we fix at 2000!
	// well we cannot place the lock here, it will make the whole thing not concurrent..
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from db is: %v \n", dbData[i])

	// m.Lock() // goroutine wait here till the lock is released, thus only 1 guy append at a time!
	// results = append(results, dbData[i])
	// m.Unlock()

	save(dbData[i])
	fetch()

	wg.Done() //decrease the counter
}

// there is chance 1 result is missing!
// The result from db is: id3
// The result from db is: id1
// The result from db is: id4
// The result from db is: id2
// total execution time: 2.001047917s
// the results are [id3 id1 id4]

// multiple go routings write into same memory location at same time!
// this is memory corruption

func save(result string) {
	m_rw.Lock() // full lock!!!, all locks must be cleared!
	results = append(results, result)
	m_rw.Unlock()
}

func fetch() {
	m_rw.RLock()
	// next read cannot go in without this read unlocks!
	// this makes the below print makes sense with the result from db is...
	// otherwise the sequence go weird
	fmt.Printf("The current results are: %v\n", results)
	m_rw.RUnlock()
}

// in this case the dbCall function is not cpu heavy
// thus if we spawn 1000 these go routine, the whole thing will still finish at about 2s
// but if the function is cpu heavy, the performance will linearly increase, based on your cpu cores
