1.go build -x -v -ldflags -s -w -buildmode=c-shared -o libhello.so github.com/ccppluagopy/ctpgo/c_call_go_so/hello

2.cd c_call_go_so
  go build -x -v -ldflags "-s -w" -buildmode=c-shared  -o libhello.so hello
  编译出 libhello.so/libhello.h

3.gcc main.c -o main -I./ -L./ -lhello 
  export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./
  ./main

