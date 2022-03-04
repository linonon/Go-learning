package twogoroutine1to100

import (
	"fmt"
	"sync"
	"testing"
)

const POOL = 100

var wg = new(sync.WaitGroup) // [Good]
var wgV2 sync.WaitGroup      // [Good]
var wgError *sync.WaitGroup  // [Error]: nil pointer

func goroutine1(ch chan int) {
	for i := 1; i <= POOL; i++ {
		ch <- i
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
	wg.Done()
}

func goroutine2(ch chan int) {
	for i := 1; i <= POOL; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	wg.Done()
}

func Test1to100WithTwoGoroutine(t *testing.T) {
	wg.Add(2)
	ch := make(chan int)
	go goroutine1(ch)
	go goroutine2(ch)
	wg.Wait()
}
