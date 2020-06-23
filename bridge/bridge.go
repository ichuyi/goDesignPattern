package bridge

import "fmt"

type Color interface {
	GetColor()string
}
type Red struct {
}

func (r *Red) GetColor()string  {
	return "red"
}

type Yellow struct {
}

func (y *Yellow)GetColor()string  {
	return "yellow"
}
type Size interface {
	GetSize()string
}
type Big struct {
}

func (b *Big)GetSize()string  {
	return "big"
}
type Small struct {
}

func (s *Small)GetSize()string {
	return "small"
}
type Bag interface {
	SetColor(c Color)
	SetSize(s Size)
	Description()
}
type MoneyBag struct {
	c Color
	s Size
}

func (m *MoneyBag)SetColor(c Color)  {
	m.c=c
}
func (m *MoneyBag) SetSize(s Size)  {
	m.s=s
}
func (m *MoneyBag)Description()  {
	fmt.Printf("this money's color is %s, size is %s\n",m.c.GetColor(),m.s.GetSize())
}
