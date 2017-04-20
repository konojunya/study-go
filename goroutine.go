package main

import (
  "fmt"
  "time"
)

func oneSecond(){
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("1 seconds passed")
}

func twoSecond(){
  time.Sleep(2000 * time.Millisecond)
  fmt.Println("2 seconds passed")
}

func threeSecond(){
  time.Sleep(3000 * time.Millisecond)
  fmt.Println("3 seconds passed")
}

func main() {
  fmt.Println(time.Now())
  go threeSecond()
  oneSecond()
  twoSecond()
  fmt.Println(time.Now())
}