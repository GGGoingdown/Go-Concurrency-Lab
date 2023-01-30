package main

import (
	"fmt"
	"time"
)

func main() {
	char := []byte("apple")
	ch1 := make(chan string)

	go func() {
		for {
			fmt.Println("Wait for data ...")
			s, ok := <-ch1
			if !ok {
				return
			}
			fmt.Println("Get data", s)
			time.Sleep(3 * time.Second)
		}
	}()

	for _, c := range char {
		ch1 <- string(c)
	}
	close(ch1)
}
