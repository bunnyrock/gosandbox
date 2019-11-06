package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = "fuck this shit"

	v := reflect.ValueOf(i)

	fmt.Printf(v.Type().Name())
}
