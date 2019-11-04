package main

import (
	"fmt"
	"sync"
	"time"
)

func runMe(){
	fmt.Println("Hello from go routine")
}

func runMe2(){
	fmt.Println("Hello from second go routine")
}

func main(){
	go runMe()
	time.Sleep(1*time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		runMe2()
		wg.Done()
	}()
	wg.Wait()

}