//nolint:unused
package main

import (
	"fmt"
	"time"
)

var cancelled = false

var mousemoveEventX = []int{}

//TODO: var events = []any //with type switch for even type

func main() {
	start := time.Now()
	fmt.Println(start.Format("05.0"), "main: started")
	busy()

	fmt.Println("main: registerMousemove")
	registerMousemove()

	// registerMouseUp()
	// registerMouseDown()
	// registerKeyPress()
	// registerAnimate()...

	fmt.Println("main: handleEvents")
	handleEvents()

	// todo: if cancelled: unregister events.
	fmt.Println(time.Now().Format("05.0"), "main finished in", time.Since(start).String())
	select {}
}

func busy() {
	fmt.Println("busy: started")
	cancelled = false
	sec := time.Now().Format("05.0")
	start := time.Now()
	for i := 1; i < 30000; i++ {
		if i%100 == 0 && time.Since(start) > time.Millisecond*25 {
			time.Sleep(time.Millisecond)
			start = time.Now()
		}
		if cancelled {
			fmt.Println("busy: cancelled")
			return
		}
		s := fmt.Sprintf("%s %s %d\n", sec, "busy", i)
		fmt.Print(s)
	}
	fmt.Println("busy: done")
}

func handleEvents() {
	cancelled = false
	for {
		time.Sleep(time.Millisecond)
		if len(mousemoveEventX) > 0 {
			lastIdx := len(mousemoveEventX) - 1
			x := mousemoveEventX[lastIdx]
			mousemoveEventX = mousemoveEventX[:lastIdx]
			fmt.Println("handleEvents: mousemove x", x)
		}
		if cancelled {
			fmt.Println("eventHandler cancelled")
			return
		}
	}
}

//export cancel
func cancel() {
	fmt.Println("cancel click")
	cancelled = true
}

//export registerMousemove
func registerMousemove()

//export mousemove
func mousemove(x, y int) {
	mousemoveEventX = append(mousemoveEventX, x)
}

/*
// TODO: implement in JS, maybe with logic to only return string after \n ...
//export jsTextInput
func jsTextInput() string

//export jsResetInput
func jsResetInput(string)

func readInput() string {
	for !cancelled {
		if s := jsTextInput(); s != "" {
			jsResetInput(s)
			return s
		}
		time.Sleep(time.Millisecond * 100)
	}
	return "" // cancelled
}
*/
