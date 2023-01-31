package main

import (
	"fmt"
	"time"
)

func main(){
	// create buffered channel
	ch := make(chan string, 10)
	// block channel
	closeChan := make(chan struct{})
	
	go func(){
		for {
			name, ok := <- ch
			if !ok{
				closeChan <- struct{}{}
				return
			}
			fmt.Println("Hello ", name)
			fmt.Printf("len=%v, cap=%v\n",len(ch), cap(ch))
			time.Sleep(1 * time.Second)
		}
	}()

	for i:=1;i<=10;i++{
		ch <- fmt.Sprintf("client #%d", i)
	}
	close(ch)

	// Wait for all done
	<-closeChan

}