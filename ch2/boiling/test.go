package main

import (
	"fmt"
	"os"
)

var a = 1

func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}
func t() error {
	if f, err := os.Open(""); err != nil { // compile error: unused: f
		return err
	} else {
		f.Close() // compile error: undefined f
	}
	return nil
}

type test struct {
	t string
}

var de = &test{}

var p = f()

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}
