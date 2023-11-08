# lesson02: Channel buffer

```shell
go mod init lesson02
```

## Channel
```go
var 채널명 chan 자료형
```

## Channel Instance 생성 + buffer
```go
var messages chan string = make(chan string, 2)
or
messages := make(chan string, 2)
```

## 예제 01 : 문제점
```go
package main

import "fmt"

func main() {
	ch := make(chan string)
	defer close(ch)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 결과
// fatal error: all goroutines are asleep - deadlock!
```

## 예제 02 : 문제점 해결
```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2)
    defer close(ch)

    ch <- "Hello"
    ch <- "World"

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}