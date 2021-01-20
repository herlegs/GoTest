package main

import (
	"code.byted.org/gopkg/pkg/testing/assert"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

}

func TestGetProductFromPSM(t *testing.T) {
	fmt.Printf("%v\n",1/float64(2))
	assert.Equal(t, "", getProductFromPSM(""))
	assert.Equal(t, "a", getProductFromPSM("a.b.c"))
	assert.Equal(t, "", getProductFromPSM("a.b"))
	assert.Equal(t, "", getProductFromPSM("a.b."))
}

