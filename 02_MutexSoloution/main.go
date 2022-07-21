package main

import (
	"fmt"
	"sync" // for using lock and unlock(mutex solution)
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()	
}

func main () {

	var w sync.WaitGroup

	// m is variable form sync.Mutex type
	var m sync.Mutex

	// This loop will create 1000 goroutines
	for i :=0; i < 1000; i++ {
		w.Add(1)
		// here we pass  both wg and m to increment function to ensoure that only one goroutine will be accessing x at a time
		//and in this case only one goroutine will be accessing x at a time and x is lock until that goroutine is done
		go increment(&w, &m)
	}
	// wait for all goroutines to finish
	w.Wait()
	fmt.Println("final value of x", x)
}