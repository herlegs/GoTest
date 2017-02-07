package reflection

import (
	"testing"
	"reflect"
	"fmt"
	"encoding/json"
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

func TestUnmarshalStruct(t *testing.T) {
	book := &book{}
	str := `{"name":"Harry potter","Content":"stir egg"}`
	UnmarshalStruct(str, book)
	fmt.Println("pl:",book)
	printStruct(book)

	gift := &gift{}
	str = `{"note":"mygift to you","name":"your life","content":"is good"}`
	UnmarshalStruct(str, gift)
	fmt.Println("pl",gift)
	printStruct(gift.book)
}

func printStruct(dto interface{}){
	bytes,err := json.Marshal(dto)
	if(err != nil){
		fmt.Println("err while marshaling:",err)
	}
	fmt.Println(string(bytes))
}

func TestReflectSetVal(t *testing.T){
	a := &struct {
		book
	}{}
	//canset := book{}
	fmt.Println(a)
	//cannotsetVal := reflect.ValueOf(cannotset)
	//cannotsetVal.Set(reflect.ValueOf("set cannot set"))
	//fmt.Println(cannotset)
	
	//cansetVal := reflect.ValueOf(canset)
	//toset := reflect.ValueOf("set can set")
	fmt.Println(reflect.ValueOf(a).Elem().Field(0).CanSet())
	str := `{"Name":"test", "Content":"also"}`
	UnmarshalStruct(str, a)
	fmt.Println(a)
	//inter := reflect.ValueOf(a).Elem().Field(0).Addr().Interface()
	//fmt.Println(reflect.ValueOf(inter).Elem().CanAddr())
}