package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("len:%v,%v\n", len(string(decoded)), string(decoded))
}

var base64Str = "eyJlbGVtZW50cyI6W3siYXR0cmlidXRlcyI6eyJlbWFpbCI6InRoZWEudGFuMDdAZ21haWwuY29tIiwiZnVsbE5hbWUiOiJNYXJpYSBSaXN0aGlhIEFwcGxlIFRhbiBDYcOxZXRlIn0sInR5cGUiOiJjdXN0b21lciIsImN1c3RvbWVyX2lkIjoiVTMxMjU0OTE5MjEifV0sInR5cGUiOiJ0cmFuc2l0aW9uIn0="
