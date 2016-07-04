//second timer

package main

import (
	"fmt"
	"time"
)

func pinger(c chan float64) {
	v := 1.0
	for i := 0; ; i++ {
		v += 1
		c <- v
	}
}

func printer(c chan float64) {
	for {
		msg := <-c
		fmt.Printf("%g\n", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan float64 = make(chan float64)

	go pinger(c)
	go printer(c)

	var input float64
	fmt.Scanln(&input)
}
