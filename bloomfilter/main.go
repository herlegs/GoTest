package main

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	filter := bloom.New(10000000, 5)

	r := filter.EstimateFalsePositiveRate(100)
	fmt.Printf("%v\n", r)

}
