package main

import (
	"fmt"
	"sync"
	"time"
	"encoding/json"
	"reflect"
)

func main(){
	test := &Payload{}
	testConvert(test)
	fmt.Println(test)
}

func testLockMain(){
	lock := &sync.RWMutex{}

	go testLock(lock, 12)

	time.Sleep(time.Second * 3)

	fmt.Println("test lock main function end")
}


func testLock(m *sync.RWMutex, id int) {
	defer fmt.Printf("go rountine %v ended\n", id)
	m.RLock()
	defer m.RUnlock()

	m.RLock()
	defer m.RUnlock()

	fmt.Printf("ID : %v\n", id)
}

func testClosureMain() {
	s := "begin"

	go func() {

		for {
			fmt.Println(s)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	time.Sleep(time.Second)

	s = "change"

	time.Sleep(time.Second)
}

func testMarshalMap(){
	m := make(map[string]interface{})
	m["a"] = "value of a"
	m["b"] = 2
	jsStr, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("err is : %v\n",err)
	}

	jsonStr := string(jsStr)

	fmt.Println(jsonStr)

	decoded := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &decoded)
	fmt.Println(decoded["b"])
}

type customMarsh struct {
	private string
	Public string
	Eb *Embed
}

//func (c *customMarsh) UnmarshalJSON(data []byte) error {
//	_ = json.Unmarshal(data, c)
//	return nil
//}

type Embed struct {
	E string
}

func (c *Embed) UnmarshalJSON(data []byte) error {
	c.E = "embeded value"
	return nil
}

// conclusion:
// If we want to customize decoding of a member field,
// to make customized UnmarshalJSON to be called, it must have a key value in the json string,
// otherwise this field will be empty (even if in unmarshal function we assign value to its fields
// and it should be a exported field (public)
// If we want to unmarshal private fields
// their parent should implement unmarshalers
func testCustomMarshal(){
	jsonStr := `{"public":"value1","private":"value2","eb":""}`
	custom := &customMarsh{}
	err := json.Unmarshal([]byte(jsonStr), custom)
	if err != nil {
		fmt.Println("err: ",err)
	}
	fmt.Println(custom.Eb)
}

type Payload struct {
	Name string
}

func testConvert(payload interface{}){
	str := `{"Name":"joe"}`
	p := &Payload{}
	_ = json.Unmarshal([]byte(str), p)

	var pI interface{}
	pI = p

	fmt.Println(pI)

	pIv := reflect.ValueOf(pI).Elem()

	reflect.ValueOf(payload).Elem().Set(pIv)
	fmt.Println(payload)

	//payv.Set(pIv.Addr())
}
