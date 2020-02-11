package inta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := ""
	var b json.RawMessage
	b = json.RawMessage(a)
	fmt.Printf("%v,%v\n", b == nil, len(b))

}

// A
func cc() <-chan struct{} {
	a := make(chan struct{})
	defer close(a)
	return a
}
