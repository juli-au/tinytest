//nolint:unused
package main

import "time"

//export hello
func hello() {
	println("hello:", time.Now().Format("15:04:05.0"))
}

//export sleep
func sleep() {
	println("sleep 1:", time.Now().Format("15:04:05.0"))
	time.Sleep(time.Second * 3)
	println("sleep 2:", time.Now().Format("15:04:05.0"))
}

func main() {
	println("main:", time.Now().Format("15:04:05.0"))
}
