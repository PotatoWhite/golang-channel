package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <- tick:
			fmt.Println("Tick")
		case <- ctx.Done():	// main 에서 cancel 호출 시 Done() 이 호출된다.
			wg.Done()
			return
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	wg.Add(1)
	go PrintEverySecond(ctx)

	wg.Wait()
}
```
