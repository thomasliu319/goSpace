package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		fmt.Println("loopFunc:", i)
		defer lock.Unlock()
	}
}
