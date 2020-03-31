export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:.

cd $GOPATH/testcgo/ctpgo/hello
g++ -shared -fPIC -o ../libhello.so hello.cpp

go build -x -v -ldflags "-s -w" -buildmode=c-shared -o libctpgo.so testcgo/ctpgo

cd $GOPATH/testcgo/ctpgo
g++ -I. -L. -lctpgo -lhello main.cpp