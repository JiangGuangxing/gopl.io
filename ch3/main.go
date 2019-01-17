package main

import "fmt"

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	var t2 test2
	t2.test.t = 4
	t2.t = 5
	fmt.Println(t2)
	t(&t2.test)
	fmt.Println(t2)
}
func t(tt *test) {
	tt.t = 6
}

type test struct {
	t int
}
type test2 struct {
	test
}

func tt(t int) *test {
	return &test{t}
}
func tt2() (t test) {
	t = test{}
	return
}

var tt3 = tt2()

func tt4(t test) test {
	return test{2}
}
