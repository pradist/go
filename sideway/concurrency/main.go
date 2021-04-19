package main

import (
	"fmt"
	"sync"
	"time"
)

func say(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello,", name, time.Now().Nanosecond())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go say("pooh", &wg)
	go say("pooh2", &wg)
	fmt.Println("done!!!", time.Now().Nanosecond())
	wg.Wait()
	fmt.Println("after wait", time.Now().Nanosecond())
}
