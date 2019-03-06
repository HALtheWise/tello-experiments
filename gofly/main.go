package main

import (
	"fmt"

	"github.com/HALtheWise/tello-experiments/gofly/ceres"
	"github.com/SMerrony/tello"
)

var _ tello.Tello
var _ ceres.StringVector

func example() int {
	// cs := C.CString("Hello from stdio\n")
	// n := C.myprint(cs)
	// C.free(unsafe.Pointer(cs))

	return int(100)
}

func main() {
	fmt.Println("Hello World", example())
}
