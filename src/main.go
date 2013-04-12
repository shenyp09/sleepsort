package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("The number of CPU cores is", runtime.NumCPU())

	tab := []int{1, 3, 0, 5, 8, 12, 23, 125, 43, 123, 6, 2133, 53, 23, 123, 512, 1235, 6134, 1}
	start := time.Now().Add(1 * time.Second)

	ch := make(chan int, 10000)
	for _, value := range tab {
		go func(val int) {
			wait := start.Sub(time.Now())
			time.Sleep(wait)
			time.Sleep(time.Duration(val) * time.Microsecond * 10)
			ch <- val
		}(value)
	}

	for _ = range tab {
		fmt.Println(<-ch)
	}
}
