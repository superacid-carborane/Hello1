//Effectively turned it into a math stepping thing, added 3d function "ponger"

package main

import (
  "fmt"
  "time"
)

func pinger(c chan float64) {
  v := 1.0
  for i := 0; ; i++ {
    v += v
    c <- v
  }
}

func ponger(c chan float64) {
  b := 3.0
  for i := 0; ; i++ {
  	b /= 2
  	c <- b 
  }
}

func printer(c chan float64) {
  for {
    msg := <- c
    fmt.Printf("\n")
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan float64 = make(chan float64)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input float64
  fmt.Scanln(&input)
}