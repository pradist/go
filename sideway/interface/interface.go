package main

import "fmt"

func say(s interface{}) {
	fmt.Printf("%T %v\n", s, s)
}

func main() {

	say(3.5)

	var i interface{}
	i = 2

	val := i.(int)

	fmt.Printf("%T %v\n", val, val)

}
