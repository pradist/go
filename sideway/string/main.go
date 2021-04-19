package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var output string
	for i := 0; i < 100000; i++ {
		//	fmt.Printf("%05d\t%v\n", i, "N")
		output += fmt.Sprintf("%05d\t%v\n", i, "N")
	}

	l, err := f.WriteString(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
