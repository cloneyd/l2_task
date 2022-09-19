package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Данный паттерн используется в случае, если подразумевается множество действий по обработке экземпляра структуры.
	Все экземпляры этой структуры обрабатываются минимум один раз.
*/

// Обрабатываемая структура
type Wallet struct {
	balance float64
}

func NewWallet(balance float64) *Wallet {
	return &Wallet{balance: balance}
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) SetBalance(val float64) {
	w.balance = val
}

type Operation interface {
	SetNext(Operation)
	Execute(*Wallet)
}

type AddOperation struct {
	amount float64
	next   Operation
}

func NewAddOperation(amount float64) *AddOperation {
	return &AddOperation{amount: amount}
}

func (ao *AddOperation) SetNext(operation Operation) {
	ao.next = operation
}

func (ao *AddOperation) Execute(wallet *Wallet) {
	wallet.SetBalance(wallet.Balance() + ao.amount)
	if ao.next != nil {
		ao.next.Execute(wallet)
	}
}

type WithdrawOperation struct {
	amount float64
	next   Operation
}

func NewWithdrawOperation(amount float64) *WithdrawOperation {
	return &WithdrawOperation{amount: amount}
}

func (wo *WithdrawOperation) SetNext(operation Operation) {
	wo.next = operation
}

func (wo *WithdrawOperation) Execute(wallet *Wallet) {
	wallet.SetBalance(wallet.Balance() - wo.amount)
	if wo.next != nil {
		wo.next.Execute(wallet)
	}
}

func main() {
	wallet := NewWallet(15)

	add1 := NewWithdrawOperation(10)
	add2 := NewAddOperation(20)
	add2.SetNext(add1)
	withdraw1 := NewAddOperation(20)
	withdraw1.SetNext(add2)
	withdraw1.Execute(wallet)

	fmt.Println(wallet.Balance())
}
