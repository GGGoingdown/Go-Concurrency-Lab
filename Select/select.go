package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch := make(chan int)
	// go func(){
	// 	time.Sleep(2 * time.Millisecond)
	// 	ch <- 1
	// }()
	close(ch)
	select {
	case <-ch:
		fmt.Println("ch close")
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 02")
	}
}
