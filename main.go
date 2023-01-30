package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {
	var bankBalance int
	var balance sync.Mutex

	incomes := []Income{
		{Source: "Main job", Amount: 100},
		{Source: "Part time job", Amount: 10},
		{Source: "Gifts", Amount: 1},
		{Source: "Investments", Amount: 20},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				tmp := bankBalance
				tmp += income.Amount
				bankBalance = tmp
				balance.Unlock()

				fmt.Printf("week %d get money %d from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Total bank balance %d\n", bankBalance)

}
