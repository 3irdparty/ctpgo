#include "hello.h"
#include "libctpgo.h"

int main() {
	//regist go func HelloGo to c
	RegistHello(HelloGo);

	//call go func HelloGo by c
	CallGoHello();

	//call c func HelloC by go
	CallCHello();
}