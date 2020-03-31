package main

//#include "hello.h"
//#cgo LDFLAGS: libhello.so
import "C"
import (
	"fmt"
)

func main() {}

//export HelloGo
func HelloGo(str string) {
	fmt.Println(str)
}

//export CallCHello
func CallCHello() {
	C.HelloC()
}
