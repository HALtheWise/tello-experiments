package main

import (
	"fmt"
)

// var _ tello.Tello

func example() int {
	// cs := C.CString("Hello from stdio\n")
	// n := C.myprint(cs)
	// C.free(unsafe.Pointer(cs))

	return int(100)
}

func main() {
	fmt.Println("Hello World", example())
}
