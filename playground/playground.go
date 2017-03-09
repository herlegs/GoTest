package main

import (
	"fmt"
	"sync"
)

func main(){
	cache := make(map[string]sync.Mutex)
	cache["a"] = sync.Mutex{}
	g,ok := cache["g"]
	fmt.Println(g,ok)
	g.Lock()
	g.Unlock()
	fmt.Println(cache["g"])
}
