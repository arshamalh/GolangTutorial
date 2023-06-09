package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	sync.Mutex
	balance float64
}

// shared resource
// race condition
// memory synchronization

func main() {
	wallet := Wallet{balance: 1000}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)
		// Depositor
		go func() {
			for i := 0; i < 20; i++ {
				wallet.Deposit(10)
			}
			wg.Done()
		}()

		// Withdrawer
		go func() {
			for i := 0; i < 20; i++ {
				wallet.Withdraw(10)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(wallet.balance)
}

func (w *Wallet) Deposit(value float64) {
	w.Lock()
	defer w.Unlock()
	w.balance += value
}

func (w *Wallet) Withdraw(value float64) {
	w.Lock()
	defer w.Unlock()
	w.balance -= value
}
