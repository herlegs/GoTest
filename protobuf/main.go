package protobuf

//go:generate protoc -I. -I$GOPATH/src --go_out=. test.proto

func dummy() {

}