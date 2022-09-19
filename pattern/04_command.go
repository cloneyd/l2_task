package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Данный паттерн используется в случае, если необходима независимость структуры-отправителя от структуры-получателя.
*/

// Интерфейс комманды
type Order interface {
	Execute()
}

// Reciever - структура-получатель
type Stock struct {
	name  string
	price float64
}

func NewStock(name string, price float64) *Stock {
	return &Stock{
		name:  name,
		price: price,
	}
}

func (s *Stock) Buy() {
	fmt.Printf("Bought %q stock for %f.\n", s.name, s.price)
}

func (s *Stock) Sell() {
	fmt.Printf("Sold %q stock for %f.\n", s.name, s.price)
}

// Commands - структуры, которые вызывают соответсвующий метод Reciever'а в функции Execute
type BuyStock struct {
	stock *Stock
}

func NewBuyStock(stock *Stock) *BuyStock {
	return &BuyStock{stock: stock}
}

func (bs *BuyStock) Execute() {
	bs.stock.Buy()
}

type SellStock struct {
	stock *Stock
}

func NewSellStock(stock *Stock) *SellStock {
	return &SellStock{stock: stock}
}

func (ss *SellStock) Execute() {
	ss.stock.Sell()
}

// Invoker - структура-отправитель
type Broker struct {
	orders []Order
}

func (b *Broker) Add(order Order) {
	b.orders = append(b.orders, order)
}

func (b *Broker) Place() {
	for _, order := range b.orders {
		order.Execute()
	}
}

func main() {
	twitterStock := NewStock("twitter", 70.0)

	buyTwitterStock := NewBuyStock(twitterStock)
	sellTwitterStock := NewSellStock(twitterStock)

	var broker Broker
	broker.Add(buyTwitterStock)
	broker.Add(sellTwitterStock)

	broker.Place()
}
