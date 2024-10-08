/*
This is the function we’ll run in every goroutine.
Sleep to simulate an expensive task.
This WaitGroup is used to wait for all the goroutines launched here to finish.
Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
Launch several goroutines and increment the WaitGroup counter for each.
Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
