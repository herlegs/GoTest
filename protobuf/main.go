package test

////go:generate protoc -I. -I$GOPATH/src --gogoproto=true --gogo_out=. test.proto

//go:generate protoc  -I. -I$GOPATH/src --gofast_out=plugins=grpc:. test.proto

//--proto_path=${PROTO_DIR} --proto_path=${PROTO_DIR}/streams/coban
