package main

/*
#cgo CFLAGS: -g -Wall

#include <stdlib.h>
#include "noros.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.Fortytwo())
}
