// Construct a pipeline that connects an arbitrary number of goroutines with
// channels. What is the maximum number of pipeline stages you can create
// without running out of memory? How long does a value take to transit the
// entire pipeline?
package main

import (
	"fmt"
	"time"
)

func main() {
	head := make(chan time.Time)
	tail := head

	for stageCount := 1; ; stageCount++ {
		go readTail(tail, stageCount)
		head <- time.Now()

		oldTail := tail
		tail = make(chan time.Time)
		go connect(oldTail, tail)
	}
}

func readTail(tail chan time.Time, stageCount int) {
	startTime := <-tail
	endTime := time.Now()
	fmt.Printf("%d\t%v\n", stageCount, endTime.Sub(startTime))
}

func connect(src, dst chan time.Time) {
	for t := range src {
		dst <- t
	}
}

// Memory usage was insignificant - but pipeline transit grew quickly
// Snipped output ->
// 100  329.803µs
// ...
// 1000  20.875636ms
// ...
// 41257 1m0.070234643s
