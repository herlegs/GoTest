package main

import (
	"github.com/herlegs/GoTest/reflection"
)

type private struct{
	Name string
}

func main(){
	dto := &private{Name:"name"}
	reflection.IsSettable(dto)
}
