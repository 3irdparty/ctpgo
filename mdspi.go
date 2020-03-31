package main

//##include "mdspi.h"
//#cgo LDFLAGS: thostmduserapi.so
//#cgo LDFLAGS: thosttraderapi.so
//#cgo LDFLAGS: libctpgo.so
import "C"
import (
	"fmt"
)

///询价通知
func OnRtnForQuoteRsp(mdspiId uint32, quoteRsp *C.CTPGO_CThostFtdcForQuoteRspField) {
	//printf("CtpMdSpi::OnRtnForQuoteRsp()\n");
}
