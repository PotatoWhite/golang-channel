# lesson03: Select 와 context

```shell
go mod init lesson04
```

# select 문법

```go
select {
case <- chan1:
	// ...
case <- chan2:
	// ...
}
```

## 예제 01 : multiple channel

```go
package main

import (
	"fmt"
	"time"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	cnt := 0
	for {
		select {
		case num := <-ch:
			fmt.Printf("Square: %d\n", num*num)
			cnt++
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)
	defer close(ch)
	defer close(quit)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	quit <- true
	wg.Wait()
}

// 결과
Square: 0
Square: 1
Square: 4
Square: 9
Square: 16
Square: 25
Square: 36
Square: 49
Square: 64
Square: 81
```

# context

작업을 취소하거나, 타임아웃을 설정하거나, 작업에 대한 데드라인을 설정할 수 있다.

## 예제 01 :작업 취소

```go
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
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go PrintEverySecond(ctx)

	time.Sleep(5 * time.Second)
	cancel()	// 작업 취소

	wg.Wait()
}
```

## 예제 02 : 타임아웃

```go
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