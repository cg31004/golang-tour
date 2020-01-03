package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"code.google.com/go-tour/reader"
)

func main() {
	fmt.Println("=======    part1   =======")
	Errorpart()
	fmt.Println("=======    part2   =======")
	Reader1()
	fmt.Println("=======    part3   =======")
	Reader2()
	fmt.Println("=======    part4   =======")
	ioReader()
}

// ===========    error part  ============
//MyFloat float64
type MyFloat float64

func (mf MyFloat) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(mf))
}

//Sqrt sqrt f
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, MyFloat(f)
	}
	return math.Sqrt(f), nil
}

//Errorpart error main
func Errorpart() {
	fmt.Println(Sqrt(4.3))
	fmt.Println(Sqrt(-5.2))
}

// ===========  Readers part  ============
func Reader1() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	fmt.Println(b)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// ===========  Readers part2  ============
type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	len := 0
	for i := range b {
		b[i] = 'A'
		len++
	}
	return len, nil
}
func Reader2() {
	reader.Validate(MyReader{})
}

// ===========  ioReader  ============
type rot13Reader struct {
	r io.Reader
}

func isAlphabet(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func rot13(c byte) byte {
	if 'a' <= c && 'c' <= 'z' {
		if c += 13; c > 'z' {
			c = c - 26
		}
	}
	if 'A' <= c && c <= 'Z' {
		if c += 13; c > 'Z' {
			c = c - 26
		}
	}
	return c
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, e := r.r.Read(b)
	for i := 0; i < n; i++ {
		if isAlphabet(b[i]) {
			b[i] = rot13(b[i])
		}
		if e == io.EOF {
			break
		}
	}
	return n, e
}

func ioReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	// fmt.Println(r)
	io.Copy(os.Stdout, &r)
}
