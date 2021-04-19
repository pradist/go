package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type errWrite struct {
	io.Writer
	err error
}

func (ew *errWrite) Write(p []byte) (n int, _ error) {
	if ew.err != nil {
		return 0, ew.err
	}
	n, ew.err = ew.Writer.Write(p)
	return n, nil
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {

	ew := &errWrite{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(w, "\r\n")

	io.Copy(w, body)

	return ew.err
}

//func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
//	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
//	if err != nil {
//		return err
//	}
//
//	for _, h := range headers {
//		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
//		if err != nil {
//			return err
//		}
//	}
//
//	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
//		return err
//	}
//
//	_, err = io.Copy(w, body)
//	return err
//}

func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)

	fmt.Printf("buf: \n%s\n", buf.String())
	fmt.Println("error:", err)
}
