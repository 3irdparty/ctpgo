cd $GOPATH/src/github.com/ccppluagopy/ctpgo

export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$GOPATH/src/github.com/ccppluagopy/ctpgo

export CPLUS_INCLUDE_PATH=$CPLUS_INCLUDE_PATH:$GOPATH/src/github.com/ccppluagopy/ctpgo/libctpc

go build -x -v -ldflags "-s -w" -buildmode=c-shared -o libctpgo.so github.com/ccppluagopy/ctpgo/libctpgo
