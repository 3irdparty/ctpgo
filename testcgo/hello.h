#ifndef _hello_h_
#define _hello_h_

#include "libctpgo.h"

#ifdef __cplusplus
extern "C" {
#endif

type void (*TypeHelloFunc)(GoString str);

void RegistHello(TypeHelloFunc hh);

void CallGoHello();

void HelloC();

#ifdef __cplusplus
}
#endif

#endif // _hello_h_
