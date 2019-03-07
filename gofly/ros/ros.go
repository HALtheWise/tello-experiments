package main

/*
#include <stdlib.h>
#include "version.h"
#cgo LDFLAGS: -L/opt/ros/melodic/lib
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.openCVVersion())
}
