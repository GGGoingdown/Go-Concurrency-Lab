package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name      string
	leftFork  int
	rightFork int
}

var persons = []Person{
	{name: "A", leftFork: 0, rightFork: 1},
	{name: "B", leftFork: 1, rightFork: 2},
	{name: "C", leftFork: 2, rightFork: 3},
	{name: "D", leftFork: 3, rightFork: 4},
	{name: "E", leftFork: 4, rightFork: 0},
}

var hunger = 3
var eatTime = 1 * time.Second
var sleepTime = 3 * time.Second

func main() {
	fmt.Println("--- Before dining ---")
	fmt.Println("---------------------")
	dine()
	fmt.Println("--- After dining ---")
	fmt.Println("--------------------")
}

func dine() {
	wg := sync.WaitGroup{}
	wg.Add(len(persons))

	seated := sync.WaitGroup{}
	seated.Add(len(persons))

	forks := make(map[int]*sync.Mutex)
	for i:=0; i<len(persons);i++{
		forks[i] = &sync.Mutex{}
	}


	for i, person := range persons{
		if i % 3 == 0{
			fmt.Printf("**** %s is late. wait for 5 second ***\n", person.name)
			time.Sleep(5*time.Second)
		}
		go eat(person, &wg, &seated, forks)
	}

	wg.Wait()
}

func eat(person Person, wg *sync.WaitGroup, seated *sync.WaitGroup, forks map[int]*sync.Mutex){
	defer wg.Done()

	fmt.Printf("[%s is on the seated]\n", person.name)
	seated.Done()

	seated.Wait()

	for i:=hunger;i>0;i--{
		// check both fork is available
		if person.leftFork > person.rightFork{
			forks[person.leftFork].Lock()
			forks[person.rightFork].Lock()
		}else{
			forks[person.rightFork].Lock()
			forks[person.leftFork].Lock()
		}
		fmt.Printf("%s is eating ...\n", person.name)
		time.Sleep(eatTime)
		forks[person.rightFork].Unlock()
		forks[person.leftFork].Unlock()
		
		fmt.Printf("%s take a rest ...\n", person.name)
		time.Sleep(sleepTime)
	}

	fmt.Printf("[%s is left the seated]\n", person.name)
}
