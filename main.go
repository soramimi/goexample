package main

/*
#cgo LDFLAGS: cat.o -lstdc++
#include <stdlib.h>
char const *bar()
{
	return "bar";
}
char *cat(char *a, char *b);
*/
import "C"
import "fmt"
import "unsafe"

func main() {
	a := "foo"
	b := C.GoString(C.bar())
	c := C.cat(C.CString(a), C.CString(b));
	d := C.GoString(c)
	C.free(unsafe.Pointer(c))
	fmt.Println(d);
}
