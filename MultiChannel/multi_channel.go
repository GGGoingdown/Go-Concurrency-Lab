package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const workerNumber int = 10

func main() {
	var wg sync.WaitGroup
	jobChan := make(chan int)
	errorChan := make(chan error)
	finishChan := make(chan struct{})

	wg.Add(workerNumber)
	for id := 1; id <= workerNumber; id++ {

		go func(jobID int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Doing job - %d\n", jobID)
			jobChan <- jobID
			if jobID > 1000 {
				errorChan <- errors.New("ops! something error")
			}

		}(id, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("Finished all jobs!!!!!")
		close(finishChan)
	}()

Loop:
	for {
		select {
		case id := <-jobChan:
			fmt.Printf("Finish job - %d\n", id)
		case err := <-errorChan:
			fmt.Println(err)
			break Loop
		case <-finishChan:
			return
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout!!")
			break Loop
		}
	}
}
