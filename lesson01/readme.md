# lesson01: Channel & Context

```shell
go mod init lesson01
```

## Channel
```go
var 채널명 chan 자료형
```

## Channel Instance 생성
```go
var messages chan string = make(chan string)
or
messages := make(chan string)
```

## Channel 에  데이터 넣기
```go
messages <- "I am a string data"
```

## Channel 에서 데이터 가져오기
```go
var message string = <- messages
```

## 예제 01
```go
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

// 결과
Square: 81
```