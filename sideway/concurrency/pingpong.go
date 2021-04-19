package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++

		fmt.Printf("%v %v \n", name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)

	go player("Pooh", table)
	go player("Nong", table)

	ball := new(Ball)
	table <- ball

	time.Sleep(2 * time.Second)
	<-table
	fmt.Println("Done!")
}
