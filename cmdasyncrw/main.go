package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const TICK time.Duration = 2

func main() {
	stop := make(chan bool)
	i := 0

	go func() {
		for {
			fmt.Printf("Ping: %d\n", i)
			time.Sleep(TICK * time.Second)
			i++
		}
	}()

	scnr := bufio.NewScanner(os.Stdin)

	go func() {
		fmt.Println("For exit type 'q'")
		for scnr.Scan() {
			npt := scnr.Text()
			if npt == "q" {
				stop <- true
				return
			}
			fmt.Println(npt)
		}
	}()

	<-stop
}
