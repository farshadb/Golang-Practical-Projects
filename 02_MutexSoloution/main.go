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
	var m sync.Mutex

	// This loop will create 1000 goroutines
	for i :=0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}