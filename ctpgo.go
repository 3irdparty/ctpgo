package ctpgo

//#include "ctpgo.h"
//#cgo LDFLAGS: thostmduserapi.so
//#cgo LDFLAGS: thosttraderapi.so
//#cgo LDFLAGS: libctpgo.so
import "C"

func Version() string {
	return C.GoString(C.ctpgo_version())
}
