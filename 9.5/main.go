// Write a program with two goroutines that send messages back and forth over
// two unbuffered channels in ping-pong fashion. HOw many communications per
// second can the program sustain?
package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan int)
	pong := make(chan int)
	done := make(chan struct{})

	pingpongCount := 2000000

	startTime := time.Now()

	go func() {
		for n := 0; n < pingpongCount; n++ {
			ping <- n
			<-pong
		}
		close(ping)
		close(done)
	}()

	go func() {
		for n := range ping {
			pong <- n
		}
		close(pong)
	}()

	<-done

	endTime := time.Now()
	dur := endTime.Sub(startTime)

	secs := float64(dur.Nanoseconds()) / 1000000000.0
	perSec := float64(pingpongCount) / secs
	fmt.Printf("%v\t%v ping-pongs\t%f per second\n", dur, pingpongCount, perSec)
}

// 1.580376441s	2000000 ping-pongs	1265521.269562 per second
