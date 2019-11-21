package main

import "fmt"

type fnv func(v ...string)

var f fnv

func main() {
	f = test
	f()
}

func test() {
	fmt.Printf("fuck")
}
