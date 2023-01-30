package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func UpdateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	words := []string{
		"A",
		"B",
		"C",
		"D",
	}

	wg.Add(len(words))
	for i, word := range words {
		go UpdateMessage(fmt.Sprintf("%d - %s", i, word), &wg)
	}
	wg.Wait()
}
