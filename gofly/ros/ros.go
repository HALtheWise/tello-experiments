package main

/*
#include <stdlib.h>
#include "version.h"
#cgo LDFLAGS: -L/opt/ros/melodic/lib
#cgo CXXFLAGS: -I/opt/ros/melodic/include
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.openCVVersion())
}
