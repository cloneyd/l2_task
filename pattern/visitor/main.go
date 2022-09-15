package visitor

import "fmt"

type Container[K comparable] interface {
	Accept(Visitor[K])
}

type Vector[K comparable] struct {
	data *[]K
	len  int
	cap  int
}

func (vec *Vector[K]) Accept(v Visitor[K]) {
	v.VisitForVector(vec)
}

type List[K comparable] struct {
	head *Node[K]
}

func (l *List[K]) Accept(v Visitor[K]) {
	v.VisitForList(l)
}

type Node[K comparable] struct {
	next *Node[K]
	val  K
}

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
