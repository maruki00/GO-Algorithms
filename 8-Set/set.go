package main

import "fmt"

type Set struct {
	elements map[any]bool
}

func NewSet(items ...any) *Set {
	s := &Set{}
	s.elements = make(map[any]bool)
	for _, item := range items {
		s.elements[item] = true
	}
	return s
}
func (s *Set) Add(items ...any) {
	for _, item := range items {
		s.elements[item] = true
	}
}

func (s *Set) Delete(item any) {
	delete(s.elements, item)
}
func (s Set) Get() []any {
	items := []any{}
	for i, _ := range s.elements {
		items = append(items, i)
	}
	return items
}

func main() {
	s := NewSet(2, 4, 5, 6, 7, 5, 6, 5, 7)
	for i, j := range s.Get() {
		fmt.Println(i, " - ", j)
	}
}
