package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	high := make(chan int)
	low := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(high)
		for i := range 10 {
			high <- i * 20
		}
	}()
	go func() {
		defer wg.Done()
		defer close(low)
		for i := range 10 {
			low <- i
		}
	}()

	for high != nil || low != nil {
		select {
		case <-ctx.Done():
			return

		case v, ok := <-high:
			fmt.Println(v, ok)

		default:

			select {
			case v, ok := <-high:
				if !ok {
					high = nil
					continue
				}
				fmt.Println(v, ok)

			case v, ok := <-low:
				if !ok {
					low = nil
					continue
				}
				fmt.Println(v, ok)
			}
		}
	}
	wg.Wait()

}
