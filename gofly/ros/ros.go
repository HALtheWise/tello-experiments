package main

//#cgo LDFLAGS: -L/opt/ros/melodic/lib

/*
#cgo CXXFLAGS: -I/opt/ros/melodic/include

#include <stdlib.h>
#include "bool.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Printf("%T\n", C.Bool_data(C.openCVVersion()))
}
