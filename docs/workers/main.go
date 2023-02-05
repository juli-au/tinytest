//nolint:unused
package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var cancelled = make(chan bool)

func main() {
	start := time.Now()
	fmt.Println(start.Format("05.0"), "main: started")
	js.Global().Set("onclick", js.FuncOf(cancel))
	go busy()

	fmt.Println(time.Now().Format("05.0"), "main finished in", time.Since(start).String())
	select {}
}

func busy() {
	fmt.Println("busy: started")
	sec := time.Now().Format("05.0")
	for i := 1; i < 30000; i++ {
		select {
		case <-cancelled:
			fmt.Println("busy: cancelled")
			return
		default:
		}
		s := fmt.Sprintf("%s %s %d\n", sec, "busy", i)
		fmt.Print(s)
	}
	fmt.Println("busy: done")
}

func cancel(_ js.Value, _ []js.Value) any {
	fmt.Println("cancel: started")
	cancelled <- true
	fmt.Println("cancel: done")
	return nil
}
