package main

type Wallet struct {
	balance float64
}

func main() {
	wallet := Wallet{balance: 1000}

}

func (w *Wallet) Deposit(value float64) {
	w.balance += 10
}

func (w *Wallet) Withdraw(value float64) {
	w.balance -= 10
}
