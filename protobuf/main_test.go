package test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	//"github.com/gogo/protobuf/proto"
	//"github.com/golang/protobuf/descriptor"
	//"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

func TestName(t *testing.T) {
	m := &DriverCoreData{}

	//_, md := descriptor.ForMessage(example)
	//
	//for _, fd := range md.GetField() {
	//	opts := fd.GetOptions()
	//	GetExten
	//}

	//rval := reflect.ValueOf(m)
	//props := proto.GetProperties(reflect.TypeOf(m).Elem())
	// ...

	_, md := descriptor.MessageDescriptorProto(m)

	fmt.Printf("%#v\n", *md.Field[1].Options)
	info, err := proto.GetExtension(md.Field[0].GetOptions(), E_MyType)
	//descriptorpb.FieldOptions
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("%v\n", *info.(*string))
}

type DD interface {
	Purge()
}

type data struct{}

func (d *data) Purge() {

}

func TestGG(t *testing.T) {
	var b interface{}
	b = &data{}
	_, ok := b.(DD)
	fmt.Printf("%v\n", ok)
}
