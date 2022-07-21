package main 

import (
	"fmt"
	"sync"
)

var x = 0
func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x += 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// ch is variable form chan bool type
	// ch capacity is 1 that means only one goroutine can access this channel at a time
	// and other goroutines will wait until this goroutine is done
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)

}