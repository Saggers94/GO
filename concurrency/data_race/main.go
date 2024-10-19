package main

import (
	"fmt"
	"sync"
	"time"
)

// data races is the hardest bug
func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	//1.
	// Mutex will make sure that the value of n is read and written by only one goroutine at a time
	// With lock and unlock method of the mutex we fix the datarace issue
	var m sync.Mutex

	// when we have two goroutines that are using the same variable and updating it than
	// we would have a problem called data race as we don't know in which order the goroutines will run
	// So, it's a data race and to resolve this we need go race detector
	for i := 0; i < gr; i++ {
		//Annonymous function
		go func() {
			time.Sleep(time.Second / 20)
			// lock method from mutex will blocked the access of the variable until we unlock it
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 20)
			m.Lock()
			// or we can use defer statement
			defer m.Unlock() // unlock will be called just before the function exits
			n--
			// m.Unlock()
			wg.Done()
		}()
	}

	// wait for go routine to terminate
	wg.Wait()

	fmt.Println("Final Value of n:", n)

	// We would need CGO_ENABLED=1 and also "gcc" which allows "Go" to talk to the "C Compiler"
	// with "go run -race main.go", you can see the data race
}
