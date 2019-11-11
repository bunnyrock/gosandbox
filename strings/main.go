package main

import "fmt"

func main() {
	for i, r := range "fuck this shit" {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(r)
	}
}
