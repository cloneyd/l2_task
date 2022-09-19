package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Паттерн Builder используется для конструирования экземпляров сложных структур с множеством полей.
	Также позволяет изменять внутреннее представление объекта.
	Помогает изолировать код, реализующий логику структуры и код, отвечающий за конструирование экземпляра структуры.
*/

import (
	"errors"
	"fmt"
)

var (
	emptyName    = errors.New("name cannot be empty")
	emptySurname = errors.New("surname cannot be empty")
	emptyEmail   = errors.New("email cannot be empty")
	emptyAge     = errors.New("age cannot be empty")
	emptyPhone   = errors.New("phone cannot be empty")
)

// Product
type User struct {
	name    string
	surname string
	email   string
	age     int
	phone   string
}

// Builder
type UserBuilder interface {
	MakeName(string) UserBuilder
	MakeSurname(string) UserBuilder
	MakeEmail(string) UserBuilder
	MakeAge(int) UserBuilder
	MakePhone(string) UserBuilder
	Build() (*User, error)
}

// ConcreteBuilder
// Реализуем конкретный builder
type SpecificUserBuilder struct {
	name    string
	surname string
	email   string
	age     int
	phone   string
}

func NewSpecificUserBuilder() *SpecificUserBuilder {
	return new(SpecificUserBuilder)
}

func (sub *SpecificUserBuilder) MakeName(name string) UserBuilder {
	sub.name = name
	return sub
}

func (sub *SpecificUserBuilder) MakeSurname(surname string) UserBuilder {
	sub.surname = surname
	return sub
}

func (sub *SpecificUserBuilder) MakeEmail(email string) UserBuilder {
	sub.email = email
	return sub
}

func (sub *SpecificUserBuilder) MakeAge(age int) UserBuilder {
	sub.age = age
	return sub
}

func (sub *SpecificUserBuilder) MakePhone(phone string) UserBuilder {
	sub.phone = phone
	return sub
}

func (sub *SpecificUserBuilder) Build() (*User, error) {
	u := new(User)
	u.name = sub.name
	u.surname = sub.surname
	u.email = sub.email
	u.age = sub.age
	u.phone = sub.phone

	// Для корректной работы реализуем валидацию полей
	// Возвращаем ошибку в случае несоответствия значения поля требуемому
	switch {
	case u.name == "":
		return nil, emptyName
	case u.surname == "":
		return nil, emptySurname
	case u.email == "":
		return nil, emptyEmail
	case u.age == 0:
		return nil, emptyAge
	case u.phone == "":
		return nil, emptySurname
	}

	return u, nil
}

// Пример использования паттерна
func main() {
	ub := NewSpecificUserBuilder()
	ub.MakeAge(15)
	user, err := ub.Build()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	ub1 := NewSpecificUserBuilder().MakeName("Nikita").MakeSurname("Abramov").MakeEmail("hello@outlook.com").MakeAge(20).MakePhone("+78005553535")
	user, err = ub1.Build()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
