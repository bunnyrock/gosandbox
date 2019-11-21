package main

import "fmt"

type testt struct {
	a, b int
}

func (t testt) GetA() int {
	return t.a
}

func (t *testt) SetA(a int) {
	t.a = a
}

func main() {
	var t testt
	t.SetA(10)
	fmt.Println(t.GetA())
}
