package main

import (
	"fmt"
	"sync"
)

func main() {
	cnter := GetCounter()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				cnter.Add()
			}
		}()
	}
	wg.Wait()

	fmt.Println(cnter.Count())
}
