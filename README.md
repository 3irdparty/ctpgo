step 1
	go get github.com/ccppluagopy/ctpgo

step 2
build cpp code
	g++ -shared -fPIC -o libctpgo.so *.c

step 3
copy so to LD_LIBRARY_PATH
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:.:$GOPATH/src/github.com/ccppluagopy/ctpgo


参考
	https://stackoverflow.com/questions/1713214/how-to-use-c-in-go
go生成动态库 
	http://blog.csdn.net/fyxichen/article/details/52495369
go生成静态库 
	http://blog.codeg.cn/2016/02/19/sharing-golang-package-to-C/

go build -x -v -ldflags "-s -w" -buildmode=c-shared -o libctpgo.so github.com/ccppluagopy/ctpgo/libctpgo
go help buildmode