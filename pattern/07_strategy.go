package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

const (
	premiumName     = "premium"
	premiumCoef     = 0.8
	premiumPlusName = "premium+"
	premiumPlusCoef = 0.6
)

// Strategy interface
type Status interface {
	NewPrice(float64) float64
}

// Context
type user struct {
	name   string
	email  string
	status Status
	cart   []float64
}

func newUser(name string, email string, status Status) *user {
	return &user{name: name, email: email, status: status}
}

func (u *user) Cart() []float64 {
	return u.cart
}

func (u *user) Buy(price float64) {
	u.cart = append(u.cart, u.status.NewPrice(price))
}

func (u *user) userStatus() Status {
	return u.status
}

// Concrete strategies
type PremiumStatus struct {
	name string
	coef float64
}

func NewPremiumStatus() *PremiumStatus {
	return &PremiumStatus{
		name: premiumName,
		coef: premiumCoef,
	}
}

func (ps *PremiumStatus) NewPrice(price float64) float64 {
	return price * ps.coef
}

type PremiumPlusStatus struct {
	name string
	coef float64
}

func NewPremiumPlusStatus() *PremiumPlusStatus {
	return &PremiumPlusStatus{
		name: premiumPlusName,
		coef: premiumPlusCoef,
	}
}

func (pps *PremiumPlusStatus) NewPrice(price float64) float64 {
	return price * pps.coef
}

func main() {
	users := []*user{
		newUser("test1", "test1@mail.ru", NewPremiumStatus()),
		newUser("test2", "test2@mail.ru", NewPremiumPlusStatus()),
	}

	cart := []float64{100, 20, 15}
	for _, u := range users {
		for _, item := range cart {
			u.Buy(item)
		}
	}

	for _, u := range users {
		fmt.Println(u.userStatus(), u.Cart())
	}
}
