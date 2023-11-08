package main

import (
    "fmt"
    "time"
    "sync"
)

func square(wg *sync.WaitGroup, ch chan int) {
    n := <-ch

    fmt.Printf("Square: %d\n", n*n)
    time.Sleep(time.Second)

    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int) // or "var ch chan int = make(chan int)"

    wg.Add(1)
    go square(&wg, ch)

    ch <- 9

    wg.Wait()
}