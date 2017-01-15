package reflection

import (
	"testing"
	"reflect"
	"fmt"
)

func TestSetStructWithMap(t *testing.T) {
	var dto = &book{}
	bookName := "harry potter"
	bookContent := "harry has long hair"
	var vars = map[string]string{
		"Name": bookName,
		"Content": bookContent,
	}
	SetStructWithMap(dto, vars)
	if(dto.Name != bookName || dto.Content != bookContent){
		t.Fail()
	}
}

func TestReflectSetVal(t *testing.T){
	a := &struct {
		Canset string
	}{"aaa"}
	//canset := book{}

	//cannotsetVal := reflect.ValueOf(cannotset)
	//cannotsetVal.Set(reflect.ValueOf("set cannot set"))
	//fmt.Println(cannotset)
	
	//cansetVal := reflect.ValueOf(canset)
	//toset := reflect.ValueOf("set can set")
	inter := reflect.ValueOf(a).Elem().Field(0).Addr().Interface()
	fmt.Println(reflect.ValueOf(inter).Elem().CanAddr())
}