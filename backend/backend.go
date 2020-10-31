package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Hello World")
	for {

	}
	wg.Wait()
}
