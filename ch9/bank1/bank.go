package main

import (
	"fmt"
	"sync"
)

type bank struct {
	amount int
}

func (b *bank) SaveMoney(amount int) {
	b.amount += amount
}

func (b *bank) GetAmount() int {
	return b.amount
}

func main() {
	for i := 0; i < 10; i++ {
		bankTest()
	}
	// bank amount:  300
	// bank amount:  300
	// bank amount:  300
	// bank amount:  100
	// bank amount:  300
	// bank amount:  300
	// bank amount:  300
	// bank amount:  300
	// bank amount:  300
	// bank amount:  300
}

func bankTest() {
	var b bank
	var gwg sync.WaitGroup
	gwg.Add(1)
	go func() {
		defer gwg.Done()
		b.SaveMoney(100)
		fmt.Println("bank amount: ", b.GetAmount())
	}()
	gwg.Add(1)
	go func() {
		defer gwg.Done()
		b.SaveMoney(200)
	}()
	gwg.Wait()
}
