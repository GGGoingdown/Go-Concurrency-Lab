package main

import (
	"fmt"
	"sync"
)

func main() {
	worker := 10000
	ch := make(chan int, worker/2)
	wg := sync.WaitGroup{}
	wg.Add(worker)

	for i := 1; i <= worker; i++ {
		go func(num int, ch chan<- int) {
			defer wg.Done()
			ch <- num
		}(i, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make([]int, 0, worker)
	for num := range ch {
		if num % 100 == 0{
			fmt.Println(num)
		}
		result = append(result, num)
	}

	fmt.Println(len(result))
}
