//nolint:unused
package main

import "time"

//export hello
func hello() {
	println("hello:", time.Now().Format("15:04:05.0"))
}

//export sleep
func sleep() {
	for i := 1; i < 10; i++ {
		println("sleep", i, ":", time.Now().Format("15:04:05.0"))
		time.Sleep(time.Second)
	}
}

func main() {
	println("main 1 :", time.Now().Format("15:04:05.0"))
	sleep()
	println("main 2 :", time.Now().Format("15:04:05.0"))
}
