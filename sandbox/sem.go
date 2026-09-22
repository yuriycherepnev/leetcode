package main

import (
	"fmt"
	"sync"
	"time"
)

func process(counter int) {
	fmt.Println("Start", counter)
	time.Sleep(1 * time.Second)
	fmt.Println("Stop", counter)
}

func main() {
	var wg sync.WaitGroup

	sem := make(chan struct{}, 3)

	counter := 1000

	for v := range counter {
		wg.Add(1)
		go func(v int) {
			sem <- struct{}{}
			defer func() {
				<-sem
				wg.Done()
			}()
			process(v)
		}(v)
	}

	wg.Wait()
}
