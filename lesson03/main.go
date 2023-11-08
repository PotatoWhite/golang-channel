package main

import (
	"fmt"
	"time"
	"sync"
)

func square(wg *sync.WaitGroup, ch <-chan int) {
	cnt := 0
	for n := range ch {
		fmt.Printf("%d Square: %d\n", cnt, n*n)
		cnt++
		time.Sleep(time.Second)
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)	// 채널을 닫아준다.
	wg.Wait()
}