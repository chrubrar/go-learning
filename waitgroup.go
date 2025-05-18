package main

import (
	"fmt"
	"sync"
)

// Create a mutex to protect access to num
//var mu sync.Mutex

func add(w *sync.WaitGroup, num *int32) {
	defer w.Done()

	// Lock before accessing num
	//mu.Lock()
	*num++
	//atomic.AddInt32(num, 1)
	// Unlock after we're done
	//mu.Unlock()
}

func main() {
	var num int32 = 0

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go add(wg, &num)
	}
	wg.Wait()
	fmt.Println(num)
}
