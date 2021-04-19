package main

import (
	"fmt"
)

//type Scanner struct {
//	br  *bufio.Reader
//	err error
//}
//
//func (s *Scanner) Err() error {
//	if s.err == io.EOF {
//		return nil
//	}
//
//	return s.err
//}
//
//func (s *Scanner) read() {
//	_, s.err = s.br.ReadString('\n')
//}
//
//func (s *Scanner) Scan() bool {
//	if s.err != nil {
//		return false
//	}
//
//	s.read()
//
//	return true
//}
//
//func NewScanner(r io.Reader) *Scanner {
//	return &Scanner{br: bufio.NewReader(r)}
//}
//
//// CountLines ...
//func CountLines(r io.Reader) (int, error) {
//	sc := NewScanner(r)
//	lines := 0
//
//	for sc.Scan() {
//		lines++
//	}
//
//	return lines, sc.Err()
//}

func main() {
	s := `please
count
me
five
line`

	//f := strings.NewReader(s)
	//n, err := CountLines(f)
	//fmt.Printf("n: %d, err: %v", n, err)
	fmt.Println(s)

}
