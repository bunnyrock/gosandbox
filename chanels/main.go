package main

import (
	"fmt"
	"sync"
	"time"
)

type safestr struct {
	str string
	mu  sync.Mutex
}

var str safestr

func main() {

	var wg sync.WaitGroup
	// async senders
	del := time.Now().UnixNano()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < 5; j++ {
				str.mu.Lock()
				str.str += fmt.Sprintf("%d hello ima sender %d\n", time.Now().UnixNano()-del, i)
				str.mu.Unlock()
				time.Sleep(time.Duration(j) * time.Millisecond * 100)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Print(str.str)
}
