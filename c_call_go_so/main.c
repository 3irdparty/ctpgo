#include <stdio.h>
#include "libhello.h"

int main() {
	GoString gostr = Hello();
	char *cstr = (char*)malloc(gostr.n+1);
	strncpy(cstr, gostr.p, gostr.n);
	printf("cstr: %s\n", cstr);
	Test();
}