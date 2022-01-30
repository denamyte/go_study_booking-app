package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
func chanDemo() {

	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}
