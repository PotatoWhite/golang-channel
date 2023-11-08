# lesson03: Wait 

```shell
go mod init lesson03
```

# for range를 이용한 channel 사용
```go
for message := range 채널명 {
	fmt.Println(message)
}
```

## 예제 01 : 문제점

```go
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
	defer close(ch)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg.Wait()
}

// 결과
0 Square: 0
1 Square: 1
2 Square: 4
3 Square: 9
4 Square: 16
5 Square: 25
6 Square: 36
7 Square: 49
8 Square: 64
9 Square: 81
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc00005a700?)
        /usr/local/go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0x0?)
        /usr/local/go/src/sync/waitgroup.go:116 +0x48
main.main()
        /home/potato/Spaces/Studyspace/potato_golang/study03/lesson03/main.go:31 +0xfc

goroutine 6 [chan receive]:
main.square(0x0?, 0x0?)
        /home/potato/Spaces/Studyspace/potato_golang/study03/lesson03/main.go:11 +0xd7
created by main.main in goroutine 1
        /home/potato/Spaces/Studyspace/potato_golang/study03/lesson03/main.go:26 +0xc7
exit status 2
```
10개의 값을 받고 나서도 대기하고 있기 때문에 deadlock이 발생한다.

## 예제 02 : 문제점 해결

```go
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
```
defer close(ch)를 사용하지 않고, close(ch)를 사용하여 채널을 닫아준다.

#