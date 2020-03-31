#include <stdio.h>
#include <string.h>
#include "hello.h"

TypeHelloFunc h = 0;

void RegistHello(TypeHelloFunc hh) {
	h = hh;
}

void CallGoHello() {
	if (h) {
		GoString s;
		s.p = "aaa";
		s.n = 4;
		h(s);
	}
}

void HelloC() {
	printf("HelloC()\n");
}
