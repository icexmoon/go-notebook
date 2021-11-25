package main

import (
	"fmt"
	"sync"
	"time"
)

type bank struct {
	amount    int
	saveChan  chan int
	getChan   chan int
	closeChan chan struct{}
}

func (b *bank) Init() {
	b.saveChan = make(chan int)
	b.getChan = make(chan int)
	b.closeChan = make(chan struct{})
}

func (b *bank) SaveMoney(amount int) {
	b.saveChan <- amount
}

func (b *bank) GetAmount() int {
	return <-b.getChan
}

func (b *bank) StartBank() {
	for {
		select {
		case amount := <-b.saveChan:
			b.amount += amount
		case b.getChan <- b.amount:
		case <-b.closeChan:
			return
		}
	}
}

func (b *bank) Close() {
	close(b.closeChan)
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
	b.Init()
	go b.StartBank()
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
	b.Close()
	time.Sleep(time.Second)
}
