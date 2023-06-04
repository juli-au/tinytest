package main

// #include <stdlib.h>
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("go main")
}

//export helloworld
func helloworld() {
	helloworldCB("Hello 🌏!")
}

//export helloworldCB
func helloworldCB(string)
