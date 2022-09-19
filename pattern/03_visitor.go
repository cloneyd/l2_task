package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Данный паттерн используется в случаях, когда нет возможности изменять код уже существующей структуры,
	но при этом есть необходимость добавить некоторый функционал
*/

import "fmt"

// Для функицонирования паттерна необходимо добавить в структуру один метод, который принимает Visitor
type Container[K comparable] interface {
	Accept(Visitor[K])
}

// Структура Vector (как в C++)
type Vector[K comparable] struct {
	data *[]K
	len  int
	cap  int
}

// Реализуем метод Accept, который будет вызывать соответствующий метод Visitor
func (vec *Vector[K]) Accept(v Visitor[K]) {
	v.VisitForVector(vec)
}

// Односвязанный список
type List[K comparable] struct {
	head *Node[K]
}

type Node[K comparable] struct {
	next *Node[K]
	val  K
}

func (l *List[K]) Accept(v Visitor[K]) {
	v.VisitForList(l)
}

// Visitor
type Visitor[K comparable] interface {
	VisitForVector(*Vector[K])
	VisitForList(*List[K])
}

type Sorter[K comparable] struct {
}

func (s *Sorter[K]) VisitForVector(v *Vector[K]) {
	//	sorting
	fmt.Println("sorting vector")
}

func (s *Sorter[K]) VisitForList(l *List[K]) {
	//	sorting
	fmt.Println("sorting list")
}
