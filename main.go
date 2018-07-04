package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))

	// http://grokbase.com/t/gg/golang-nuts/151tjs2gs3/go-nuts-dont-understand-cgo-printf-doesnt-print-a-thing
	C.fflush(C.stdout)
}


func main() {
	Print("Hello, cgo")
}
