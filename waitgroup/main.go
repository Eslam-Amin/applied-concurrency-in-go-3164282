package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(1) // increment the inner counter
	go hello(&wg)
	wg.Wait() // wait until the inner counter becomes 0
	goodbye()
	fmt.Println("End")
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the inner counter
	fmt.Println(("Hello, World!"))
}

func goodbye() {
	fmt.Println(("goodbye, World!"))
}
