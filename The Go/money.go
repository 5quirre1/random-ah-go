package main

import (
	"fmt"
  "math/rand"
)

func main() {
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("I am rich asf, what to buy greg??????")
  } else {
    fmt.Println("Where did all my money go vro ?!?!?!?!?!?!?!?!?")
  }
}
