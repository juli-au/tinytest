//nolint:unused
package main

import (
	"fmt"
	"time"
)

var cancelled = false
var ch = make(chan bool)

func main() {
	fmt.Println("go:main: started")

	go pingpong()
	busy()

	fmt.Println("go:main: done")
}

func busy() {
	fmt.Println("busy: started")
	cancelled = false
	start := time.Now()
	for i := 1; i < 30000; i++ {
		if time.Since(start) > time.Millisecond*25 {
			fmt.Println("=========================== yield ===")
			// ch <- true // doesn't work
			// runtime.Gosched() // doesn't work.
			time.Sleep(time.Millisecond) // works
			start = time.Now()
		}
		if cancelled {
			fmt.Println("go:busy: cancelled")
			return
		}
		fmt.Printf("%s %d\n", "go:busy", i)
	}
	fmt.Println("go:busy: done")
}

func pingpong() {
	for range ch {
	}
}

//export cancel
func cancel() {
	fmt.Println("go: cancel")
	cancelled = true
}
