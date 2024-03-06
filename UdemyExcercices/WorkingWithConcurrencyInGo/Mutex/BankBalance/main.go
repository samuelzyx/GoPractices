package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func ByIncome(incomes []Income, balanceBank *int, balanceMutex *sync.Mutex) {
	for i, income := range incomes {
		ByWeek(i, income, balanceBank, balanceMutex)
	}
}

func ByWeek(i int, income Income, balanceBank *int, balanceMutex *sync.Mutex) {
	defer wg.Done()
	for week := 1; week <= 52; week++ {
		balanceMutex.Lock()
		temp := *balanceBank
		temp += income.Amount
		*balanceBank = temp
		balanceMutex.Unlock()

		fmt.Printf("On week %d, you earned $%d.00 from %s, Total Balance: %d\n", week, income.Amount, income.Source, *balanceBank)
	}
}

func main() {
	var balanceBank int
	var balanceMutex sync.Mutex

	fmt.Printf("Initial account balance: $%d.00\n", balanceBank)

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg.Add(len(incomes))

	ByIncome(incomes, &balanceBank, &balanceMutex)

	wg.Wait()

	fmt.Printf("Final bank balance: $%d.00\n", balanceBank)
}
