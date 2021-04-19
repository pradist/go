package main

import (
	"fmt"
	"time"
)

func say2(name string, c chan int) {
	fmt.Println("Hello,", name, time.Now().Nanosecond())
	c <- 1
}

func main() {

	c := make(chan int)

	go say2("Pooh", c)
	fmt.Println("Wait")
	<-c
	fmt.Printf("c:%v\n", c)
	fmt.Println("Done!")
}
