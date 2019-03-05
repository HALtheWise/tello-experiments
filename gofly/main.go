package main

/*
#include <stdio.h>
#include <stdlib.h>

int myprint(char* s) {
	printf("%s\n", s);
	return 1;
}
*/
import "C"

import "fmt"
import "unsafe"

import "github.com/SMerrony/tello"

var _ tello.Tello

func example() int {
	cs := C.CString("Hello from stdio\n")
	n := C.myprint(cs)
	C.free(unsafe.Pointer(cs))

	return int(n)
}

func main() {
	fmt.Println("Hello World", example())
}
