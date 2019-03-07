package main

/*
#include <stdlib.h>
#include "version.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.openCVVersion())
}
