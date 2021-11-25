package main

import (
	"fmt"
	"sync"
)

type bank struct {
	amountMutex sync.RWMutex
	amount      int
}

func (b *bank) SaveMoney(amount int) {
	b.amountMutex.Lock()
	b.saveMoney(amount)
	b.amountMutex.Unlock()
}

func (b *bank) saveMoney(amount int) {
	b.amount += amount
}

func (b *bank) GetAmount() int {
	b.amountMutex.RLock()
	defer b.amountMutex.RUnlock()
	return b.getAmount()
}

func (b *bank) getAmount() int {
	return b.amount
}

func (b *bank) WithDraw(amount int) bool {
	b.amountMutex.Lock()
	defer b.amountMutex.Unlock()
	b.saveMoney(-amount)
	if b.getAmount() < 0 {
		b.saveMoney(amount)
		return false
	}
	return true
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
		b.WithDraw(50)
		fmt.Println("bank amount: ", b.GetAmount())
	}()
	gwg.Add(1)
	go func() {
		defer gwg.Done()
		b.SaveMoney(200)
	}()
	gwg.Wait()
}
