package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Scanner struct {
	br *bufio.Reader
	err error
}

func (sc *Scanner) Err() error {
	if sc.err == io.EOF {
		return nil
	}
	return sc.err
}

func (sc *Scanner) Scan() bool {
	if sc.err != nil {
		return false
	}
	sc.Read()
	return true
}

func (sc *Scanner) Read() {
	_, sc.err = sc.br.ReadString('\n')
}

func countLines(r io.Reader) (int, error) {

	sc := Scanner {br: bufio.NewReader(r)}
	lines := 0

	for sc.Scan() {
		lines++
	}

	//for {
	//	_, err = br.ReadString('\n')
	//	lines++
	//	if err != nil {
	//		break
	//	}
	//}
	//
	//if err != io.EOF {
	//	return 0, err
	//}

	return lines, nil
}

func main() {
	s := `please
count
me
five
line`

	f := strings.NewReader(s)
	n, err := countLines(f)
	fmt.Printf("n: %d, err: %v", n, err)
}
