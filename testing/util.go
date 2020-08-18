package testing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	defaultTimeOut = time.Second * 5
)

//HTTPGet is helper function of sending http get request with params
func HTTPGet(url string, dto interface{}, header interface{}) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("err make new http get request")
		return nil, err
	}
	addParamToRequest(request, dto)
	addHeaderToRequest(request, header)
	client := http.Client{
		Timeout: defaultTimeOut,
	}
	return client.Do(request)
}

//DecodeFromResponse unmarshal resposne body into obj
func DecodeFromResponse(resp *http.Response, obj interface{}) {
	err := json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		log.Println(err)
	}
}

func addParamToRequest(request *http.Request, dto interface{}) {
	values := request.URL.Query()
	jsonBytes, err := json.Marshal(dto)
	if err != nil {
		log.Println("err marshal params")
		return
	}
	objMap := map[string]*json.RawMessage{}
	err = json.Unmarshal(jsonBytes, &objMap)
	if err != nil {
		log.Println(err)
		return
	}
	for key, val := range objMap {
		//remove quotes
		bytes, _ := unquoteBytes(*val)
		values.Add(key, string(bytes))
	}
	request.URL.RawQuery = values.Encode()
}

func addHeaderToRequest(request *http.Request, header interface{}) {

}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	return s[1 : len(s)-1], true
}

//PrintObj prints any struct object
func PrintObj(obj interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("err while marshaling:", err)
	}
	fmt.Println(string(bytes))
}
